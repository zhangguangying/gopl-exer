package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := make([]string, 0)
	for i := 0; i < 1000; i++ {
		s = append(s, "hello world")
	}
	handJoinTime := time.Now()
	handStr := ""
	for i := 0; i < len(s); i++ {
		handStr += s[i]
	}
	fmt.Printf("hand join string used time: %v\n", time.Since(handJoinTime).Microseconds())

	stringJoinTime := time.Now()
	strings.Join(s, "")
	fmt.Printf("string join string used time: %v\n", time.Since(stringJoinTime).Microseconds())
}
