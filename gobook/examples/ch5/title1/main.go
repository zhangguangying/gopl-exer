package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		resp.Body.Close()
	}()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

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

func main() {
	title(os.Args[1])
}