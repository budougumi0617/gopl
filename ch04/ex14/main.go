package main

import (
	"log"
	"net/http"
	"os"
)

// GitHubHandler shows GitHub repository information.
type GitHubHandler struct {
	store *LocalStore
}

// RepoHandler renders repository information.
func (h *GitHubHandler) RepoHandler(w http.ResponseWriter, r *http.Request) {
	h.store.RenderIssues(w)
	h.store.RenderMilestones(w)
	h.store.RenderUsers(w)
}

func main() {
	url := "golang/go" // Default URL

	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	handler := &GitHubHandler{NewLocalStore()}
	handler.store.Load(url)
	log.Println("Access localhost:8080/" + url)
	http.HandleFunc("/"+url, handler.RepoHandler)
	http.ListenAndServe(":8080", nil)
}
