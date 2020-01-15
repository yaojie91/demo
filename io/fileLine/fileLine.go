package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var line int

	for {
		_, isPrefix, err := r.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			line++
		}
	}

	fmt.Println(line)
}
