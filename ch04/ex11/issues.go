package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/seikichi/gopl/ch04/ex11/github"
)

var usage = `Usage issues <command> [options]

Commands:
  create
  show
  search
  edit
  close
  reopen

Environment Variables:
  GITHUB_ACCESS_TOKEN: github access token
                       (see https://github.com/settings/tokens)
`

func getStringByEditor() string {
	f, err := ioutil.TempFile("/tmp", "github-issues")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("vi", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func main() {
	token := os.Getenv("GITHUB_ACCESS_TOKEN")
	client := github.NewClient(token)

	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	switch args[0] {
	case "create":
		create(client, args[1:])
	case "show":
		show(client, args[1:])
	case "search":
		search(client, args[1:])
	case "edit":
		edit(client, args[1:])
	case "close":
		close(client, args[1:])
	case "reopen":
		reopen(client, args[1:])
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func getOwnerAndRepo(fs *flag.FlagSet, args []string) (string, string, []string) {
	if len(args) < 2 {
		fs.Usage()
		os.Exit(1)
	}
	s := strings.Split(args[0], "/")
	return s[0], s[1], args[1:]
}

func getNumber(fs *flag.FlagSet, args []string) (string, []string) {
	if len(args) < 1 {
		fs.Usage()
		os.Exit(1)
	}
	return args[0], args[1:]
}

func printIssue(issue *github.Issue) {
	fmt.Printf("Number:    %d\n", issue.Number)
	fmt.Printf("Title:     %s\n", issue.Title)
	fmt.Printf("State:     %s\n", issue.State)
	fmt.Printf("Body:      %q\n", issue.Body)
	fmt.Printf("URL:       %s\n", issue.HTMLURL)
	fmt.Printf("CreatedAt: %s\n", issue.CreatedAt)
}

func printIssueList(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func search(client *github.Client, args []string) {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(`Usage issues search <owner>/<repo>`)
	}
	owner, repo, args := getOwnerAndRepo(fs, args)
	issues, err := client.SearchIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	printIssueList(issues)
}

func show(client *github.Client, args []string) {
	fs := flag.NewFlagSet("show", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(`Usage issues show <owner>/<repo> <number>`)
	}
	owner, repo, args := getOwnerAndRepo(fs, args)
	number, args := getNumber(fs, args)
	issue, err := client.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}

func create(client *github.Client, args []string) {
	fs := flag.NewFlagSet("edit", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Print(`Usage issues create <owner>/<repo> [options]

Options:
  --title
  --body (set EDITOR if you want to use editor)
`)
	}

	var title, body string
	fs.StringVar(&title, "t", "", "title")
	fs.StringVar(&title, "title", "", "title")
	fs.StringVar(&body, "b", "", "body")
	fs.StringVar(&body, "body", "", "body")

	owner, repo, args := getOwnerAndRepo(fs, args)
	fs.Parse(args)

	if body == "EDITOR" {
		body = getStringByEditor()
	}

	issue, err := client.CreateIssue(owner, repo, &github.IssueCreateParams{
		Title: title,
		Body:  body,
	})
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}

func edit(client *github.Client, args []string) {
	fs := flag.NewFlagSet("edit", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Print(`Usage issues edit <owner>/<repo> <number> [options]

Options:
  --title
  --body (set EDITOR if you want to use editor)
`)
	}

	var title, body string
	fs.StringVar(&title, "t", "", "title")
	fs.StringVar(&title, "title", "", "title")
	fs.StringVar(&body, "b", "", "body")
	fs.StringVar(&body, "body", "", "body")

	owner, repo, args := getOwnerAndRepo(fs, args)
	number, args := getNumber(fs, args)
	fs.Parse(args)

	if body == "EDITOR" {
		body = getStringByEditor()
	}

	issue, err := client.EditIssue(owner, repo, number, &github.IssueEditParams{
		Title: title,
		Body:  body,
	})
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}

func close(client *github.Client, args []string) {
	fs := flag.NewFlagSet("close", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(`Usage issues close <owner>/<repo> <number>`)
	}
	owner, repo, args := getOwnerAndRepo(fs, args)
	number, args := getNumber(fs, args)
	issue, err := client.EditIssue(owner, repo, number, &github.IssueEditParams{State: "closed"})
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}

func reopen(client *github.Client, args []string) {
	fs := flag.NewFlagSet("close", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(`Usage issues reopen <owner>/<repo> <number>`)
	}
	owner, repo, args := getOwnerAndRepo(fs, args)
	number, args := getNumber(fs, args)
	issue, err := client.EditIssue(owner, repo, number, &github.IssueEditParams{State: "open"})
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}
