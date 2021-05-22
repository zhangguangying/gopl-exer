package main

import (
	"fmt"
	"testing"
)

func TestSearchIssues(t *testing.T) {
	s := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, _ := SearchIssues(s)
	for _, item := range result.Items {
		fmt.Println(item.User)
	}
}
