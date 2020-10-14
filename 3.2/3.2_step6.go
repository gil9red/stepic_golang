package main

import "fmt"

func main() {
	text := "строка"

	bs := []byte(text)
	fmt.Println(text, string(bs))
	fmt.Println(bs) // [209 129 209 130 209 128 208 190 208 186 208 176]

	fmt.Println()

	rs := []rune(text)
	fmt.Println(text, string(rs))
	fmt.Println(rs) // [1089 1090 1088 1086 1082 1072]
}
