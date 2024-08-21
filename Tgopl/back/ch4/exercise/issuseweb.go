package main

import (
	"hello/github"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	var result *github.IssuesSearchResult
	var keywords = []string{"php", "js"}
	result, _ = github.SearchIssue(keywords)

	var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	  <th>#</th>
	  <th>State</th>
	  <th>User</th>
	  <th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
	  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	  <td>{{.State}}</td>
	  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))
	issueList.Execute(w, result)
}
