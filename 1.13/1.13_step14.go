package main

import (
	"fmt"
	"log"
)

/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи
оно является, то есть выведите такое число n, что φn=A. Если А не является
числом Фибоначчи, выведите число -1.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
8

Sample Output:
6
*/

func fib(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, b+a
	}
	return a
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < n; i++ {
		if n == fib(i) {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}
