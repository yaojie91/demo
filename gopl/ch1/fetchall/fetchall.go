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
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("total time: %.2fs", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ch <- fmt.Sprintf("%.2fs  %d  %s", time.Since(start).Seconds(), n, url)
}
