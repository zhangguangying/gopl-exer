package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	stack := outline(nil, doc)
	fmt.Println("stack")
	fmt.Println(stack)
}

func outline(stack []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
	return stack
}
