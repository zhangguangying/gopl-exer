package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"testing"
)

func TestHtml(t *testing.T) {
	resp, _ := http.Get("https://zhangguangying.cn")
	r := html.NewTokenizer(resp.Body)
	for {
		t := r.Next()
		if t == html.ErrorToken {
			break
		}
		fmt.Printf("t: %s\n", )
		/*fmt.Printf("text: %s\n", r.Text())
		tagName, _ := r.TagName()
		fmt.Printf("tagname: %s\n", tagName)
		fmt.Printf("token: %s\n", r.Token())
		fmt.Printf("raw: %s\n", r.Raw())*/

	}
}
