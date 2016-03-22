package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	month := []*github.Issue{}
	year := []*github.Issue{}
	other := []*github.Issue{}

	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)

	for _, is := range result.Items {
		if is.CreatedAt.After(monthAgo) {
			month = append(month, is)
		} else if is.CreatedAt.After(yearAgo) {
			year = append(year, is)
		} else {
			other = append(other, is)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("\n%d issues (less than a month old):\n", len(month))
	printIssues(month)
	fmt.Printf("\n%d issues (less than a year old):\n", len(year))
	printIssues(year)
	fmt.Printf("\n%d issues (more than a year old):\n", len(other))
	printIssues(other)
}

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006/01/02"), item.User.Login, item.Title)
	}
}
