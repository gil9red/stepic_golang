package main

import (
	"fmt"
)

/*
Напишите функцию, находящую наименьшее из четырех введённых в
этой же функции чисел.

Входные данные
Вводится четыре числа.

Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7

Sample Output:
4
*/

func minimumFromFour() int {
	var min int
	_, _ = fmt.Scan(&min)

	for i := 1; i < 4; i++ {
		var x int
		_, _ = fmt.Scan(&x)

		if x < min {
			min = x
		}
	}

	return min
}

func main() {
	fmt.Println(minimumFromFour())
}
