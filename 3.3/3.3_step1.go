package main

import (
	"fmt"
	"strings"
	"unicode"
)

func invert(r rune) rune {
	// Если буква строчная, то она возвращается заглавной
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	// Иначе возвращается строчной
	return unicode.ToLower(r)
}

func returnFunction() func(rune) rune {
	return invert
}

func main() {
	src := "aBcDeFg"
	test := "AbCdEfG"

	// Обратите внимание, что скобки после имени функции используются только при ее вызове
	src1 := strings.Map(invert, src)
	src2 := strings.Map(returnFunction(), src)

	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src1 == test)
	// Инвертированная строка: AbCdEfG. Результат: true.

	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src2 == test)
	// Инвертированная строка: AbCdEfG. Результат: true.

	fmt.Println()

	fmt.Println(src == strings.Map(invert, strings.Map(invert, src)))
	// true
}
