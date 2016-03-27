package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	e = nil
	source, _ := ioutil.ReadFile("./ex08.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	n := ElementByID(doc, "sameid")
	for _, a := range n.Attr {
		if a.Key == "href" && a.Val != "http://first" {
			t.Errorf("Not first element: %s", a.Val)
		}
	}

}

func TestMain(t *testing.T) {
	e = nil
	os.Args = []string{"", "https://github.com/budougumi0617/404status", "NotExistID____"}
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if got != "Not found id element\n" {
		t.Errorf("got:\n%s", got)
	}

	e = nil
	os.Args = []string{"", "https://golang.org", "menu"}
	stdout = new(bytes.Buffer) // captured output
	main()
	got = stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("Not output\n")
	}
}
