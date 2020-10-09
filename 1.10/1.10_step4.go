package main

import "fmt"

/*
Напишите программу, которая в последовательности чисел находит сумму двузначных
чисел, кратных 8. Программа в первой строке получает на вход число n - количество
чисел в последовательности, во второй строке -- n чисел, входящих в данную
последовательность.

Sample Input:
5
38 24 800 8 16

Sample Output:
40
*/

func main() {
	var n, num int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&num)

		if num >= 10 && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}
