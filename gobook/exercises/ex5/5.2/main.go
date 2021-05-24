package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, _ := html.Parse(os.Stdin)
	m := make(map[string]bool)
	count := 0
	ElementCount(&count, m, doc)
	fmt.Println(count)
}

func ElementCount(count *int, m map[string]bool, n *html.Node) {
	if n.Type == html.ElementNode && !m[n.Data] {
		m[n.Data] = true
		*count++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ElementCount(count, m, c)
	}
}
