package main

import (
	"fmt"
	"gopl-exer/gobook/examples/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "%#-5d %9.9s %.55s\n"

	now := time.Now()
	pastDay := make([]*github.Issue, 0)
	pastMon := make([]*github.Issue, 0)
	pastYea := make([]*github.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0,0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreatedAt.After(month):
			pastMon = append(pastMon, item)
		case item.CreatedAt.After(year):
			pastYea = append(pastYea, item)
		}
	}

	if len(pastDay) > 0 {
		fmt.Printf("\nPast Day:\n")
		for _, item := range pastDay {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastMon) > 0 {
		fmt.Printf("\nPast Month:\n")
		for _, item := range pastMon {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastYea) > 0 {
		fmt.Printf("\nPast Year:\n")
		for _, item := range pastYea {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
}