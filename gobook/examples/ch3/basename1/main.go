package main

import "fmt"

func main() {
	s1 := "a.go"
	s2 := "a/b/c.go"
	s3 := "a/b.c.go"
	fmt.Println(basename(s1))
	fmt.Println(basename(s2))
	fmt.Println(basename(s3))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i + 1:]
			break;
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
