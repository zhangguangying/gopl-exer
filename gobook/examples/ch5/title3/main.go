package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func setTitle() (err error) {
	defer func() {
		if p := recover(); p != nil {
			if p == "A" {
				err = fmt.Errorf("error string")
			}
		}
	}()
	err = fmt.Errorf("error another")
	panic("A")
	return err
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
	err := setTitle()
	if err != nil {
		fmt.Println(err)
	}
}
