package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int
func startElement(n *html.Node) {
	fmt.Printf("<%s>\n",n.Data)
}
func endElement(n *html.Node) {
	fmt.Printf("<%s>\n", n.Data)
}

func main() {
	resp, _ := http.Get(os.Args[1])
	node, _ := html.Parse(resp.Body)
	forEachNode(node, startElement, endElement)
}
