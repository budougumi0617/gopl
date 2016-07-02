// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templ = template.Must(template.New("itemlist").Parse(`
  <!DOCTYPE html>
<html>
<head>
<title>ch07ex12</title>
<style>
table {
	border-collapse: collapse;
}
td, th {
	border: solid 1px;
	padding: 0.5em;
  text-align: right;
}
</style>
</head>
<body>
  <table>
  <tr><th>Name</th><th>Price</th></tr>
    {{range .}}
    <tr>
      <td>{{.Name}}</td><td>{{.Price | printf "%4.2f"}}</td>
    </tr>
    {{end}}
  </table>
</body>
</html>`))

// Item name and price
type Item struct {
	Name  string
	Price dollars
}

func main() {
	db := database{"shoes": 50, "socks": 5, "desk": 300, "note": 120}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	items := []Item{}
	for n, p := range db {
		items = append(items, Item{n, p})
	}
	if err := templ.Execute(w, items); err != nil {
		log.Fatal(err)
	}
}
