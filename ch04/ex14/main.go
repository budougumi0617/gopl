package main

import "log"

func main() {
	ls := NewLocalStore()
	if err := ls.Load("golang/go"); err != nil {
		log.Fatal(err)
	}
	log.Println("test!!")
}
