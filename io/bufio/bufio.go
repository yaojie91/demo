package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	strReader := strings.NewReader("hello world")
	bufReaer := bufio.NewReader(strReader)
	data, _ := bufReaer.Peek(5)
	fmt.Println("-----------", data, bufReaer.Buffered())

	str, _ := bufReaer.ReadString(' ')
	fmt.Println("--------_", str, bufReaer.Buffered())

	w := bufio.NewWriter(os.Stdout)
	n, _ := fmt.Fprint(w, "hello")
	err := w.Flush()
	if err != nil {
		return
	}
	fmt.Println("--------", n)
}
