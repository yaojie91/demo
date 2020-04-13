package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	start := time.Now()
	ch := make(chan string)
	url := os.Args[1]
	//for _, url := range os.Args[1:] {
	go func() {
		fetch(url, ch)
		fmt.Println("1 done")
		//wg.Done()
	}()
	go func() {
		fetch(url, ch)
		fmt.Println("2 done")
		//wg.Done()
	}()
	//}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//wg.Wait()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	fineName := fileName() + ".txt"
	file, err := os.Create(fineName)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func fileName() string {
	str := "1234567890abcdef"
	sl := []byte(str)
	result := []byte{}
	for i := 0; i < 3; i++ {
		result = append(result, sl[rand.Intn(16)])

	}
	return string(result)
}
