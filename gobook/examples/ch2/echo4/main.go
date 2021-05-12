package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "seprator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Printf("%t\n", *n)
	if !*n {
		fmt.Println()
	}
}
