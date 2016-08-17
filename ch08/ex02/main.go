// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"flag"
	"log"
	"net"
)

var op string

func initialize() {
	flag.StringVar(&op, "port", "8000", "Port number") // Get -port option
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	initialize()
	initCmdMap()
	listener, err := net.Listen("tcp", "localhost:"+op)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		log.Printf("Connect from %v\n", conn.RemoteAddr())
		fh := newFtpHandler(conn)
		go fh.execute()

	}
}
