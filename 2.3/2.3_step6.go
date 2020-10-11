package main

import "fmt"

/*
Поменяйте местами значения переменных на которые ссылаются указатели.
После этого переменные нужно вывести.

Sample Input:
2 4

Sample Output:
4 2
*/

func test(x1, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func main() {
	a, b := 5, 4
	test(&a, &b)
}
