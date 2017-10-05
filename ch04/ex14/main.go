package main

import (
	"log"
	"os"
)

func main() {
	ls := NewLocalStore()
	if err := ls.Load("golang/go"); err != nil {
		log.Fatal(err)
	}
	ls.RenderIssues(os.Stdout)
	ls.RenderUsers(os.Stdout)
	ls.RenderMilestones(os.Stdout)
	log.Println("test!!")
}
