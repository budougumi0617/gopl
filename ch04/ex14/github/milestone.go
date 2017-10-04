// Copyright 2017 budougumi0617 All Rights Reserved.

package github

import (
	"fmt"
	"html/template"
	"io"
)

// Milestone is milestone for a repository.
// https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository
type Milestone struct {
	Title       string
	State       string
	Description string
}

type Milestones struct {
	Items []*Milestone
}

func (m *Milestones) RenderUsers(w io.Writer) {
	milestoneList := template.Must(template.New("milestoneList").Parse(`
	<h1>Milestones</h1>
	<table>
	  <th>Title</th><th>Description</>
	  {{range .Items}}
	  <tr>
	    <td>{{.Title}}</td><td>{{.Description}}</td>
	  </tr>
	  {{end}}
	</table>
	`))

	if err := milestoneList.Execute(w, &m); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
