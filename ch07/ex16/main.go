// Copyright 2017 budougumi0617 All Rights Reserved.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	handler := &MyForm{Answer: ""}
	log.Println("Access localhost:8080/")
	http.HandleFunc("/", handler.RenderForm)
	http.ListenAndServe(":8080", nil)
}

type MyForm struct {
	Answer string
}

// RenderForm writes simple form.
func (m *MyForm) RenderForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Request %s\n", r.Method)
		r.ParseForm()
		f := r.Form["formula"]
		log.Printf("Request formula : %s\n", f)
		m.Answer = "10"
	}
	form := template.Must(template.New("formTmpl").Parse(`
	<html>
	<head>
	  <title>Calculator</title>
	</head>
	<body>
	  <form action="/" method="post">
	    Formula : <input type="text" name="formula">
	    <input type="submit" value="enter">
	  </form>
	  <p>{{.Answer}}</p>
	</body>
	</html>
	`))

	if err := form.Execute(w, &m); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
