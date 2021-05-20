package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("C:\\Users\\hyuii\\code\\go\\src\\gopl-exer\\gobook\\exercises\\ex4\\4.9\\4.9.go")
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Printf("line\tcount\n")
	for k, v := range counts {
		fmt.Printf("%q\t%d\n", k, v)
	}
}
