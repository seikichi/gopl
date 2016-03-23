package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/seikichi/gopl/ch04/ex14/github"
)

// か き か け

type RepositoryInfo struct {
	Issues       []*github.Issue
	Milestones   []*github.Milestone
	Contributors []*github.User
}

type RepositoryInfoCache map[string]RepositoryInfo

func main() {
	token := os.Getenv("GITHUB_ACCESS_TOKEN")
	client := github.NewClient(token)

	port := 8080
	flag.IntVar(&port, "p", 8000, "listening port")
	flag.Parse()

	cache := RepositoryInfoCache{}

	pageTemplate := template.Must(template.New("repository").Parse(`
<html>
<body>
<h1>Bug Reports</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Issues}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>

<h1>Milestones</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Title</th>
</tr>
{{range .Milestones}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>

<h1>Contributors</h1>
<table>
<tr style='text-align: left'>
  <th>Login</th>
</tr>
{{range .Contributors}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
</tr>
{{end}}
</table>
</body>
</html>
`))

	handler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		owner := query.Get("owner")
		repo := query.Get("repo")

		info := cache.get(client, owner, repo)
		pageTemplate.Execute(w, info)
	}
	http.HandleFunc("/", handler)
	log.Printf("Serving lissajous on localhost port %d ...\n", port)
	log.Println("Supported queries: owner, repo")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func (c *RepositoryInfoCache) get(client *github.Client, owner, repo string) RepositoryInfo {
	if info, ok := (*c)[owner+"/"+repo]; ok {
		return info
	}

	issues, _ := client.ListIssues(owner, repo)
	milestones, _ := client.ListMilestones(owner, repo)
	users, _ := client.ListContributors(owner, repo)

	info := RepositoryInfo{
		Issues:       issues,
		Milestones:   milestones,
		Contributors: users,
	}
	(*c)[owner+"/"+repo] = info
	return info
}
