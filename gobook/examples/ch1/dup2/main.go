package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) > 0 {
		for _, filename := range os.Args[1:] {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1:%v\n", err);
				continue;
			}
			countLines(f, counts)
			f.Close()
		}
	} else {
		countLines(os.Stdin, counts)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(f.Name())
		}
	}
}
