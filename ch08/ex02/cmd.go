// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	initCmdMap()
	cmd := cmdMap["tess"]
	if cmd != nil {
		cmd(nil, "foooo")
	} else {
		fmt.Println("cmd is not found")
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var cmdMap = map[string]func(net.Conn, string){}

// initCmdMap Set function to each command string
func initCmdMap() {
	cmdMap["USER"] = notImpl
	cmdMap["QUIT"] = notImpl
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

func notImpl(con net.Conn, parm string) {
	fmt.Println(parm)
}
