package main

import "fmt"

func uniq(s []string) []string {
	i := 0;
	for j := 1; j < len(s); j++ {
		if s[j] != s[j-1] {
			s[i] = s[j]
			i++
		}
	}
	return s[:i]
}

func main() {
	s := []string{"hello", "hello", "world", "world", "shijie"}
	fmt.Println(uniq(s))
}
