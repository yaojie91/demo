package main

import (
	"fmt"
	"unicode"
)

func main() {
	// unicode包
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))
	single = '\ufe35'
	fmt.Println(unicode.IsControl(single))
	fmt.Printf("%c", single)
	fmt.Println()

	digit := '1'
	fmt.Println(unicode.IsDigit(digit))
	fmt.Println(unicode.IsNumber(digit))

	letter := 'Ⅷ'
	fmt.Println(unicode.IsDigit(letter))
	fmt.Println(unicode.IsNumber(letter))

	han := 'f'
	fmt.Println(unicode.IsDigit(han))
	fmt.Println(unicode.Is(unicode.Han, han))
	fmt.Println(unicode.In(han, unicode.Gujarati, unicode.White_Space))
}
