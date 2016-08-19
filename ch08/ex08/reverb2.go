// Copyright 2016 budougumi0617 All Rights Reserved.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// Limit is timeout seconds to client.
var Limit time.Duration = 10

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	s := make(chan string)
	wg := sync.WaitGroup{}
	// Exit for loop if client that shouts nothing within Limit seconds.
	func() {
		for {
			// Need to scan on another channel.
			go func() {
				if input.Scan() {
					s <- input.Text()
				}
			}()
			select {
			case text := <-s:
				wg.Add(1)
				go func() {
					echo(c, text, 1*time.Second)
					wg.Done()
				}()
			case <-time.After(Limit * time.Second):
				log.Println("Time out")
				return // Exit for loop.
			}
		}
	}()
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
