package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")

	ReadExample()
}

func ReadExample() {
	for {
		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err error
		)
		dir, _ := os.Getwd()

		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入不多于9个字符，以回车结束：")
			data, err = ReadFrom(os.Stdin, 10)
		case "2":
			file, err := os.Open(dir + "/../../test.go")
			if err != nil {
				fmt.Println("err is: ", err)
			}
			data, err = ReadFrom(file, 10)
			file.Close()
		case "3":
			data, err = ReadFrom(strings.NewReader("abcdefghhh"), 3)
		case "4":
			fmt.Println("暂未实现")
		case "q":
			os.Exit(0)
		}
		if err != nil {
			return
		}
		fmt.Println("-----------" ,data)
	}
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if err != nil {
		return p, err
	}
	return p[:n], nil
}