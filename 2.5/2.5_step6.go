package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Для получения количества символов используйте utf8.RuneCountInString()
*/

func main() {
	var en = "english"
	var ru = "русский"
	fmt.Println(len(en), len(ru))
	// 7 14

	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
	// 7 7
}
