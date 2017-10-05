// Copyright 2017 budougumi0617 All Rights Reserved.

package github

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// User user information
// https://developer.github.com/v3/users/#get-all-users
// https://api.github.com/users
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Users struct {
	Link  string
	Items []*User
}

func (u *Users) GetUsers() err {
	req, _ := c.newRequest(context.Background(), "GET", GitHubAPIURL+"users", nil)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status is not ok: %v", resp.StatusCode)
	}

	if err := decodeBody(resp, &(i.Items)); err != nil {
		return err
	}
	return nil
}

func (u *Users) RenderUsers(w io.Writer) {
	userList := template.Must(template.New("userList").Parse(`
	<h1>GitHub Users</h1>
	<table>
	  <th>Name</th><th>URL</>
	  {{range .Items}}
	  <tr>
	    <td>{{.Login}}</td><td>{{.HTMLURL}}</td>
	  </tr>
	  {{end}}
	</table>
	`))

	if err := userList.Execute(w, &u); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
