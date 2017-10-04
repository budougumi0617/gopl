// Copyright 2017 budougumi0617 All Rights Reserved.

package github

import (
	"html/template"
	"io"
)

// User user information
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Users struct {
	Link  string
	Items []*User
}

func (u *Users) RenderUsers(w io.Writer) {
	userList := template.Must(template.New("user").Parse(`
	<h1>GitHub Users</h1>
	<table>
	  <th>Name</th><th>URL</>
	  {{range .Users}}
	  <tr>
	    <td>{{.Login}}</td><td>{{.HTMLURL}}</td>
	  </tr>
	  {{end}}
	</table>
	`))

}