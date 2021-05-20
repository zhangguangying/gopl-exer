package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	seed := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if !seed[input.Text()] {
			seed[input.Text()] = true
			fmt.Println(input.Text())
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err);
		os.Exit(1)
	}
}
