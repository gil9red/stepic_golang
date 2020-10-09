package main

import (
	"fmt"
)

/*
На ввод подаются пять чисел, которые записываются в массив.
Вам нужно написать фрагмент кода, с помощью которого можно найти
и вывести максимальное число в этом массиве.

Sample Input:
2
3
56
45
21

Sample Output:
56
*/

func main() {
	var n uint
	fmt.Scan(&n)

	items := make([]int, n)
	for i := range items {
		fmt.Scan(&items[i])
	}

	maxN := items[0]
	for i := 1; i < len(items); i++ {
		x := items[i]
		if x > maxN {
			maxN = x
		}
	}

	fmt.Println(maxN)
}
