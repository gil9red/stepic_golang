package main

import "fmt"

/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих
направлениях (например, "топот", "око", "заказ").

Sample Input:
топот

Sample Output:
Палиндром
*/

func ToReversed(text string) string {
	r := []rune(text)

	reversed_bs := make([]rune, len(r))
	for i := 0; i < len(reversed_bs); i++ {
		reversed_bs[i] = r[len(reversed_bs)-i-1]
	}
	return string(reversed_bs)
}

func main() {
	var text string
	_, _ = fmt.Scan(&text)

	reversed := ToReversed(text)

	if text == reversed {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
