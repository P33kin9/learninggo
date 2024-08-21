package main

import (
	"hello/github"
	"html/template"
	"log"
	"os"
	"time"
)

//!+template
const templ = `{{.TotalCount}} issues;
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

//!-template

//!+daysAgo
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//!-dayAgo

//!+exec
var report = template.Must(template.New("issueslist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func main() {
	//!+parse
	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	//	noMust()
}

//!-exec

func noMust() {
	//!+parse
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	//!-parse
	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
