// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

type clock struct {
	name, hostport string
	conn           net.Conn
	time           chan string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(stderr, "Needs input \"NAME=HOST:PORT\" at least one.")
		return
	}
	var clocks []*clock
	var labels string
	for _, arg := range os.Args[1:] {
		data := strings.Split(arg, "=")
		clocks = append(clocks, &clock{name: data[0], hostport: data[1], conn: nil, time: make(chan string)})
	}
	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.hostport)
		c.conn = conn
		if err != nil {
			log.Fatal(err)
		}
		labels += c.name + "\t"
		go updateTime(c)
	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(stdout, "%s\n", labels)
		for _, c := range clocks {
			// Print each result
			fmt.Fprintf(stdout, "%s\t", <-c.time)
		}
		fmt.Fprintln(stdout, "")
	}
}

func updateTime(c *clock) {
	defer c.conn.Close()
	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		c.time <- s.Text()
	}
}
