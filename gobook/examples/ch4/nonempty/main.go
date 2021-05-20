package main

import "fmt"

func main() {
	s := []string{"Hello", "", "World"}
	//fmt.Println(nonempty(s))
	fmt.Println(noempty2(s))

	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
}

func nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func noempty2(s []string) []string {
	out := s[:0]
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}