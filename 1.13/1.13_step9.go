package main

import (
	"fmt"
)

/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество минимальных элементов.

Sample Input:
3
21
11
4

Sample Output:
1
*/

func main() {
	var n, x int
	fmt.Scan(&n)

	fmt.Scan(&x)
	min := x
	number := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x < min {
			min = x
			number = 0
		}
		if x == min {
			number++
		}
	}
	fmt.Println(number)
}
