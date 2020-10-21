package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	src := "aBcDeFg"
	test := "AbCdEfG"

	// Обратите внимание, что скобки после имени функции используются только при ее вызове
	src = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, src)

	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)
	// Инвертированная строка: AbCdEfG. Результат: true.

	fmt.Println()

	// Присваивание переменной значение анонимной функции
	fn := func(a, b int) int { return a + b }
	fmt.Println(fn(17, 15))
	// 32

	invertBin := func(r rune) rune {
		switch r {
		case '0':
			return '1'
		case '1':
			return '0'
		default:
			return r
		}
	}
	fmt.Println(strings.Map(invertBin, "bin: 101010"))
	// bin: 010101

	// Выполняем анонимную функцию на месте
	// Обратите внимание на использование скобок при вызове функции
	func(a, b int) {
		fmt.Println(a + b)
	}(12, 34)
	// 46
}
