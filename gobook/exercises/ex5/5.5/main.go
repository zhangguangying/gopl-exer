package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)


func main() {
	url := os.Args[1]
	words, images, _ := CountWordsAndImages(url)
	fmt.Printf("words:%d\timages:%d\n", words, images)
}


func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return 
	}
	words, images = countWordsAndImages(doc)
	return 
}

func countWordsAndImages(n *html.Node) (words, images int) {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return
}

func wordCount(data string) int {
	words := 0
	input := bufio.NewScanner(strings.NewReader(data))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}
	return words
}
