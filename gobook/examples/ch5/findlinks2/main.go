package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

const usage = `
usage
findlinks url
`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}
	url := os.Args[1]
	links, _ := findlinks(url)
	for _, link := range links {
		fmt.Println(link)
	}
}

func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("findlinks: get %s error:%v", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("findlinks: get %s error:%v", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("findlinks: parse error:%v", err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}