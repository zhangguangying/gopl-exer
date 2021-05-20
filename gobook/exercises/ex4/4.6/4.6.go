package main

import (
	"fmt"
	"unicode"
)

func mergeSpace(b []byte) []byte {
	s := []rune(string(b))
	j := 0
	for i := 0; i < len(s); {
		if unicode.IsSpace(s[i]) {
			s[j] = ' '
			for unicode.IsSpace(s[i]) && i < len(s) {
				i++
			}
		} else {
			s[j] = s[i]
			i++
		}
		j++
	}
	return []byte(string(s[:j]))
}

func main() {
	s := []byte("hello world     	nihao		ma")
	fmt.Printf("%s\n", mergeSpace(s))
}