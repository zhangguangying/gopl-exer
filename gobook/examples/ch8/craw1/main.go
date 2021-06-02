package main

import "os"

func main() {
	worklist := make(chan []string)

	go func() {worklist <- os.Args[1:]}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
