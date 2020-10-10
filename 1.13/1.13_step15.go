package main

import (
	"fmt"
	"log"
)

/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные
Задано единственное число N

Выходные данные
Необходимо вывести требуемое представление числа N.

Sample Input:
12

Sample Output:
1100
*/

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%b", n)
}
