package main

import (
	"fmt"
	"gopl-exer/gobook/examples/ch5/links"
	"log"
	"os"
)

func breathFirst(f func(item string) []string, worklist []string)  {
	seen := make(map[string]bool)
	if len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breathFirst(crawl, os.Args[1:])
}