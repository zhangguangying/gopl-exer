package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchAll(url, ch)
	}
	for i := 0; i < len(os.Args[1:]); i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("elapsed: %v", time.Since(start).Seconds())
}

func fetchAll(uri string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(uri)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		ch <- fmt.Sprintf("fetchall:%v", err);
		return
	}
	f, err := os.Create(url.QueryEscape(uri))
	fmt.Println(url.QueryEscape(uri))
	if err != nil {
		ch <- fmt.Sprintf("fetchall:create file:%v", err)
		return;
	}
	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetchall reading %s: %v", uri, err);
		return;
	}
	ch <- fmt.Sprintf("%.2f	%7d	 %s", time.Since(start).Seconds(), nbytes, uri)
}

