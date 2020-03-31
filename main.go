package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func SampleReadStdin() {
	fmt.Println("please enter:")
	data, err := ReadFrom(os.Stdin, 10)
	if err != nil {
		return
	}
	fmt.Println("---------", data)
}

func main() {
	SampleReadStdin()
}
