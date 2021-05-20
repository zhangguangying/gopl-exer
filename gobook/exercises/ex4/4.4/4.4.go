package main

import "fmt"

func rotate(s []int, x int) []int {
	st := make([]int, len(s), cap(s))
	copy(st, s)
	for i := 0; i < len(s); i++ {
		st[(i+x)%len(s)] = s[i]
	}
	return st
}

func main() {
	s := []int{1,2,3,4,5}
	fmt.Println(rotate(s, 2))
}
