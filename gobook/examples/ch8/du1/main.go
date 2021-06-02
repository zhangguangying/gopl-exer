package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	flag.Parse()
	roots := flag.Args()
	fmt.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	n.Add(1)
	go func() {
		for _, root := range roots {
			walkDir(root, &n, fileSizes)
		}

	}()
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {

			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
