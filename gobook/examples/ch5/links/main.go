package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
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

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %v",url, err)
	}
	links := make([]string, 0)
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val);
					if err != nil {
						continue
					}
					links = append(links, link.String())
				}
			}
		}
	}, nil)
	return links, nil
}
