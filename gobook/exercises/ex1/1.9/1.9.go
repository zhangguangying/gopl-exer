package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		defer func() {
			resp.Body.Close()
		}()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
	}
}
