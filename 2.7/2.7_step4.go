package main

import (
	"fmt"
)

/*
Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает
1000 знаков и строка содержит только десятичные цифры.

Выходные данные
Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:
1112221112

Sample Output:
2
*/

func main() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid input values")
	}

	bs := []rune(text)
	maxR := bs[0]

	for i := 1; i < len(bs); i++ {
		r := bs[i]
		if r > maxR {
			maxR = r
		}
	}

	fmt.Println(string(maxR))
}
