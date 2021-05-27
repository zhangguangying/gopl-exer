package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)=%d\n", x, x-0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x-1)
}

func printStack() {
	s := make([]byte, 0)
	n := runtime.Stack(s, false)
	os.Stdout.Write(s[:n])
}

func main() {
	f(3)
}