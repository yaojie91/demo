package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "sdfasfafa"
	fmt.Println(s)
	fmt.Println(&s)
	tmp := []byte(s)
	fmt.Println(tmp, &tmp[0], &tmp[1], &tmp[3])
	s = "qweqweqweqweqwe"
	fmt.Println(s)
	fmt.Println(&s)
	runess := '好'
	fmt.Println(utf8.ValidRune(runess))

	str := "你好"
	fmt.Println(utf8.ValidString(str))

	by := []byte("街")
	fmt.Println(by[:1], by[:2], by[:3])
	fmt.Println(utf8.Valid(by[:3]))

	s3 := "china中国"
	fmt.Println([]byte(s3))
	fmt.Println(utf8.RuneLen('\u0014'))
	fmt.Println(utf8.RuneCount([]byte(s3)))
	fmt.Println(utf8.RuneCountInString(s3))
	fmt.Println("-----", utf8.RuneLen([]rune(s3)[0]))

	p := make([]byte, 3)
	utf8.EncodeRune(p, '好')
	fmt.Println(p)
	fmt.Println(utf8.DecodeRune(p))
	fmt.Println(utf8.DecodeRuneInString("你好"))
	fmt.Println(utf8.DecodeLastRune([]byte("你好")))
	fmt.Println(utf8.DecodeLastRuneInString("你好"))

	fmt.Println(utf8.FullRune([]byte(s3)[:1]))
	fmt.Println(utf8.FullRuneInString("你好"))

	fmt.Println("11111111111")
	word := []byte("界")
	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))
}
