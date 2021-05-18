package main

import "fmt"

func main()  {
	fmt.Println(comma("1234567"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}
