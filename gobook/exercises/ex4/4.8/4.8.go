package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := map[string]int{
		"letter": 0,
		"number": 0,
		"others": 0,
	}
	invalid := 0
	input := bufio.NewReader(os.Stdin)
	for  {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsNumber(r) {
			counts["number"]++
		} else {
			counts["others"]++
		}
	}
	fmt.Printf("cate\tcount\n")
	for k, v := range counts {
		fmt.Printf("%q\t%d\n", k, v)
	}
}

