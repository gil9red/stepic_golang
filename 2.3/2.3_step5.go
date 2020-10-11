package main

import "fmt"

/*
Напишите функцию, которая умножает значения на которые ссылаются два указателя
и выводит получившееся произведение в консоль.

Sample Input:
2 2

Sample Output:
4
*/

func test(x1, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func main() {
	a, b := 5, 4
	test(&a, &b)
}
