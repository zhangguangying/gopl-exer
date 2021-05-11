package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := make([]string, 10)
	for i := 0; i < len(s); i++ {
		s = append(s, "hello")
	}
	handJoinTime := time.Now()
	handStr := ""
	for i := 0; i < len(s); i++ {
		handStr += s[i]
	}
	fmt.Printf("hand join string used time: %v\n", time.Since(handJoinTime).Seconds())

	stringJoinTime := time.Now()
	strings.Join(s, "")
	fmt.Printf("string join string used time: %v\n", time.Since(stringJoinTime).Seconds())
}
