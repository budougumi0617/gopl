// Copyright 2017 budougumi0617 All Rights Reserved.

package github

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// Milestone is milestone for a repository.
// https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository
// https://api.github.com/repos/golang/go/milestones
type Milestone struct {
	Title       string
	State       string
	Description string
}

type Milestones struct {
	Items []*Milestone
}

func (m *Milestones) GetMilestones() err {
	req, _ := c.newRequest(context.Background(), "GET", GitHubAPIURL+url+"/milestones", nil)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status is not ok: %v", resp.StatusCode)
	}

	if err := decodeBody(resp, &(m.Items)); err != nil {
		return err
	}
	return nil
}

func (m *Milestones) RenderMilestones(w io.Writer) {
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
