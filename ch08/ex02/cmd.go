// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strings"
)

// define Method Expression.
type ftpCommand func(*ftpHandler, ...string) string

// each client information.
type ftpHandler struct {
	conn            net.Conn
	addr            string
	precmd          string
	passiveListener net.Listener
	path            string
}

// root path of FTP server.
var rootpath = "./"

// map of implemantation command.
var cmdMap = map[string]ftpCommand{}

// return FTP client
func newFtpHandler(con net.Conn) *ftpHandler {
	return &ftpHandler{con, con.RemoteAddr().String(), "", nil, rootpath}
}

// start ftp connection.
func (fh *ftpHandler) execute() {
	// Reply succeed messages to client.
	fh.sendmsg("220 Service ready for new user")
	defer fh.conn.Close()
	s := bufio.NewScanner(fh.conn)
	for s.Scan() {
		inputs := fh.parsetext(s.Text())
		resp := fh.handle(inputs)
		fh.sendmsg(resp)
	}
}

func (fh *ftpHandler) parsetext(t string) []string {
	log.Printf("Parse text : %q\n", t)
	return strings.Split(t, " ")
}

// execute FTP command.
func (fh *ftpHandler) handle(msg []string) (rsp string) {
	cmd := fh.getcmd(msg[0])
	if len(msg) < 2 {
		rsp = cmd(fh, "")
	} else {
		rsp = cmd(fh, msg[1:]...)
	}
	fh.precmd = strings.ToUpper(msg[0])
	return
}

// get from cmdMap
func (fh *ftpHandler) getcmd(t string) (cmd ftpCommand) {
	t = strings.ToUpper(t)
	log.Printf("Search command : %s\n", t)
	if c, exist := cmdMap[t]; !exist {
		cmd = (*ftpHandler).notImpl
	} else {
		cmd = c
	}
	return
}

func parseAddr(address string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(address, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2), nil
}

// initCmdMap sets Method Expression to each command.
func initCmdMap() {
	cmdMap["USER"] = (*ftpHandler).handleUSER
	cmdMap["QUIT"] = (*ftpHandler).handleQUIT
	cmdMap["PWD"] = (*ftpHandler).handlePWD
	cmdMap["PORT"] = (*ftpHandler).handlePORT
	cmdMap["TYPE"] = (*ftpHandler).handleTYPE
	cmdMap["STRU"] = (*ftpHandler).handleSTRU
	cmdMap["RETR"] = (*ftpHandler).handleRETR
	cmdMap["STOR"] = (*ftpHandler).handleSTOR
	cmdMap["NOOP"] = (*ftpHandler).handleNOOP
	cmdMap["CWD"] = (*ftpHandler).handleCWD
	cmdMap["LIST"] = (*ftpHandler).handleLIST
}

// send message to client.
func (fh *ftpHandler) sendmsg(s string) {
	log.Printf("Response: %s\n", s)
	io.WriteString(fh.conn, s+"\r\n")
}

// return another connection
func (fh *ftpHandler) dataConn() (conn io.ReadWriteCloser, err error) {
	switch fh.precmd {
	case "PORT":
		conn, err = net.Dial("tcp", fh.addr)
		if err != nil {
			return nil, err
		}
	case "PASV":
		conn, err = fh.passiveListener.Accept()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported previous command")
	}
	return conn, nil
}

func (fh *ftpHandler) notImpl(parms ...string) string {
	return "502 Command not implemented."
}

// Each below comments are corresponding command and replied codes.

// USER
//  230
//  530
//  500, 501, 421
//  331, 332
func (fh *ftpHandler) handleUSER(parms ...string) string {
	return "230 User logged in, proceed."
}

// QUIT
//    221
//    500
func (fh *ftpHandler) handleQUIT(parms ...string) string {
	return "221 Service closing control connection. Logged out if appropriate."
}

// PWD
//    257
//    500, 501, 502, 421, 550
func (fh *ftpHandler) handlePWD(parms ...string) string {
	log.Printf("Execute PWD %q\n", fh.conn.RemoteAddr())
	pathname := path.Join(rootpath, fh.path)
	return "257 \"" + pathname + "\" is current directory."
}

// PORT
//    200Ø
//    500, 501, 421, 530
func (fh *ftpHandler) handlePORT(parms ...string) string {
	if len(parms) != 1 {
		return "501 Usage: PORT a,b,c,d,p1,p2"
	}
	var err error
	fh.addr, err = parseAddr(parms[0])
	if err != nil {
		log.Printf("%v", err)
		return "501 Can't parse address."
	}
	return "200 PORT command oky."
}

// TYPE
//    200
//    500, 501, 504, 421, 530
func (fh *ftpHandler) handleTYPE(parms ...string) string {
	return "502 Command not implemented"
}

// STRU
//    200
//    500, 501, 504, 421, 530
func (fh *ftpHandler) handleSTRU(parms ...string) string {
	if len(parms) != 1 {
		return "501 STRU Syntax error."
	}
	if parms[0] != "F" {
		return "504 Command not implemented for that parameter."

	}
	return "200 STRU command okay."
}

// RETR
//    125, 150
//       (110)
//       226, 250
//       425, 426, 451
//    450, 550
//    500, 501, 421, 530
func (fh *ftpHandler) handleRETR(parms ...string) string {
	if len(parms) != 1 {
		return "501 RETR Syntax error."
	}
	filename := path.Join(fh.path, parms[0])
	if _, err := os.Stat(filename); err != nil {
		log.Printf("%v", err)
		return "550 File not found."
	}
	fh.sendmsg("150 File status okay; about to open data connection.")
	conn, err := fh.dataConn()
	if err != nil {
		return "425 Can't open data connection"
	}
	defer conn.Close()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("%v\n", err)
		return "550 File unavailable."
	}

	if _, err = conn.Write(b); err != nil {
		log.Printf("%v\n", err)
		return "426 Connection closed, transfer aborted."
	}

	return "250 Requested file action okay, completed."

}

// STOR
//    125, 150
//       (110)
//       226, 250
//       425, 426, 451, 551, 552
//    532, 450, 452, 553
//    500, 501, 421, 530
func (fh *ftpHandler) handleSTOR(parms ...string) string {
	if len(parms) != 1 {
		return "501 STOR Syntax error in parameters"
	}
	filename := path.Join(fh.path, parms[0])
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("%v\n", err)
		return "550 File can't be created."
	}
	defer file.Close()
	fh.sendmsg("150 Ok to send data.")
	conn, err := fh.dataConn()
	if err != nil {
		return "425 Can't open data connection"
	}
	defer conn.Close()
	_, err = io.Copy(file, conn)
	if err != nil {
		log.Printf("%v\n", err)
		return "450 File unavailable."
	}
	return "226 Closing data connection. Requested file action successful. " + file.Name()
}

// NOOP
// 200
// 500 421
func (fh *ftpHandler) handleNOOP(parms ...string) string {
	return "200 Command okay."
}

// CWD
//    250
//    500, 501, 502, 421, 530, 550
func (fh *ftpHandler) handleCWD(parms ...string) string {
	if len(parms) != 1 {
		return "501 CWD Syntax error in parameters"
	}
	dirpath := path.Join(fh.path, parms[0])

	info, err := os.Stat(dirpath)
	if err != nil || !info.IsDir() {
		return "550 Directory not found."
	}
	fh.path = dirpath
	return "250 Requested file action okay, completed. Move to " + dirpath
}

// LIST
//    125, 150
//       226, 250
//       425, 426, 451
//    450
//    500, 501, 502, 421, 530
func (fh *ftpHandler) handleLIST(parms ...string) string {
	var filename string
	switch len(parms) {
	case 0:
		filename = fh.path
	case 1:
		filename = path.Join(fh.path, parms[0])
	default:
		return "501 Too many arguments."
	}
	file, err := os.Open(filename)
	if err != nil {
		return "550 File not found."
	}
	fh.sendmsg("150 Here comes the directory listing.")
	w, err := fh.dataConn()
	if err != nil {
		return "425 Can't open data connection."
	}
	defer w.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Printf("%v", err)
		return "450 Requested file action not taken. File unavailable."
	}
	if stat.IsDir() {
		filenames, derr := file.Readdirnames(0)
		if derr != nil {
			return "550 Can't read directory."
		}
		for _, f := range filenames {
			_, err = fmt.Fprintf(w, "%s\r\n", f)
			if err != nil {
				log.Printf("%v", err)
				return "426 Connection closed; transfer aborted."
			}
		}
	} else {
		_, err = fmt.Fprintf(w, "%s\r\n", filename)
		if err != nil {
			log.Printf("%v\n", err)
			return "426 Connection closed; transfer aborted."
		}
	}
	return "226 Closing data connection. Requested file action successful"
}
