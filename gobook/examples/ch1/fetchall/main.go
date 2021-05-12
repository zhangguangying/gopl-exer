package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

func fetchAll(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		ch <- fmt.Sprintf("fetchall:%v", err);
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetchall reading %s: %v", url, err);
		return;
	}
	ch <- fmt.Sprintf("%.2f	%7d	 %s", time.Since(start).Seconds(), nbytes, url)
}
