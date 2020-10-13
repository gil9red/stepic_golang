package main

import (
	"fmt"
	"strconv"
)

/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа
и вывести получившееся число.

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1.
Единица в квадрате - 1. В итоге получаем 811181

Sample Input:
9119

Sample Output:
811181
*/

func main() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid input values")
	}

	for _, r := range text {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Sprintf("invalid int value: %c", r))
		}

		fmt.Print(n * n)
	}

	fmt.Println()
}
