// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type ftpHandler struct {
	conn net.Conn
}

func newFtpHandler(con net.Conn) *ftpHandler {
	return &ftpHandler{con}
}

var cmdMap = map[string]func(...string) string{}

func (fh *ftpHandler) execute() {
	// Reply succeed messages to client.
	fh.sendmsg("220 Service ready for new user")
	defer fh.conn.Close()
	s := bufio.NewScanner(fh.conn)
	for s.Scan() {
		inputs := fh.parsetext(s.Text())
		cmd := fh.getcmd(inputs[0])
		var msg string
		if len(inputs) < 2 {
			msg = cmd("")
		} else {
			msg = cmd(inputs[1:]...)
		}
		fh.sendmsg(msg)
	}
}

func (fh *ftpHandler) parsetext(t string) []string {
	log.Printf("Parse text : %q\n", t)
	return strings.Split(t, " ")
}

func (fh *ftpHandler) getcmd(t string) (cmd func(...string) string) {
	t = strings.ToUpper(t)
	log.Printf("Search command : %s\n", t)
	if c, exist := cmdMap[t]; !exist {
		cmd = notImpl
	} else {
		cmd = c
	}
	return
}

// initCmdMap Set function to each command string
func initCmdMap() {
	cmdMap["USER"] = handleUSER
	cmdMap["QUIT"] = handleQUIT
	cmdMap["PWD"] = handlePWD
	cmdMap["PORT"] = notImpl
	cmdMap["TYPE"] = notImpl
	cmdMap["MODE"] = notImpl
	cmdMap["STRU"] = notImpl
	cmdMap["RETR"] = notImpl
	cmdMap["STOR"] = notImpl
	cmdMap["NOOP"] = notImpl
	cmdMap["CMD"] = notImpl
	cmdMap["LIST"] = notImpl
}

func (fh *ftpHandler) sendmsg(s string) {
	log.Printf("Response: %s\n", s)
	io.WriteString(fh.conn, s+"\r\n")
}

func notImpl(parms ...string) string {
	return "202 Command not implemented, superfluous at this site."
}

func handleUSER(parms ...string) string {
	return "230 User logged in, proceed."
}

func handleQUIT(parms ...string) string {
	return "221 Service closing control connection. Logged out if appropriate."
}

func handlePWD(parms ...string) string {
	pathname := "./"
	return "257 \"" + pathname + "\" is current directory."
}
