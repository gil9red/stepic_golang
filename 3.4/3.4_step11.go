package main

import (
	"fmt"
	"unicode"
)

type customError uint

func (c customError) Error() string {
	return fmt.Sprintf("digit!!!: index %d", c)
}

func errorInString(str string) error {
	// Полезная работа со строкой проигнорирована
	for i, r := range []rune(str) {
		if unicode.IsDigit(r) {
			return customError(i)
		}
	}
	return nil
}

func main() {
	err := errorInString("Строка1Строка")
	if err != nil {
		fmt.Printf("Ошибка обработана: %v\n", err)
	}
	// Ошибка обработана: digit!!!: index 6

	if idx, ok := err.(customError); ok {
		fmt.Printf("Контекст: %d\n", idx)
	}
	// Контекст: 6
}
