// 模拟读取、解析、写入log的过程
package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	dbsource string
}

type LogProcess struct {
	rc    chan string
	wc    chan string
	read  Reader
	write Writer
}

func (rf *ReadFromFile) Read(rc chan string) {
	msg := "message"
	rc <- msg
}

func (wt *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println("------", <-wc)
}

func (lp *LogProcess) Process() {
	data := <-lp.rc
	lp.wc <- strings.ToUpper(data)
}
func main() {
	r := &ReadFromFile{path: "/tmp/1.log"}
	w := &WriteToInfluxDB{dbsource: "fasf"}
	lp := &LogProcess{
		make(chan string),
		make(chan string),
		r,
		w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1 * time.Second)
}
