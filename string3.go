package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "gopher"
	b := "hello world"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("壹", "一"))

	s := "测试包含con"
	fmt.Println(strings.Contains(s, "one"))
	fmt.Println(strings.ContainsAny(s, "测量"))
	fmt.Println(strings.ContainsRune(s, 'o'))
	fmt.Println(strings.Index(s, "on"))

	fmt.Println(strings.Count("fivevev", "vev"))
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	fmt.Println(strings.SplitAfter("abcde", ""))

	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasSuffix("Gopher", "r"))

	han := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println("--------”", strings.IndexFunc("fsdsdfdf,世界", han))
	fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))

	mapping := func(r rune) rune {
		return '六'
	}
	fmt.Println(strings.Map(mapping, "6666asfa66"))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 5))
	fmt.Println(strings.ToLower("EKEKEKFE"))
	fmt.Println(strings.ToUpper("feeiifeef"))

	x := "!!!@@@你好,!@#$ Gophe你好rs###$$$"
	f := func(r rune) bool {
		return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	}
	fmt.Println("------", strings.TrimFunc(x, f))
	fmt.Println(strings.TrimLeftFunc(x, f))
	fmt.Println(strings.TrimRightFunc(x, f))
}
