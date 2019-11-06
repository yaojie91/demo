package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("你好"))
	////fmt.Println("-------", b)
	//fmt.Fprint(&b, ",", "http://www.flysnow.org")
	////把Buffer里的内容打印到终端控制台
	//b.WriteTo(os.Stdout)

	var p [100]byte
	n, err := b.Read(p[:])
	if err != nil {
		fmt.Println("--------", err)
	}
	fmt.Println(n, err, string(p[:n]))
}
