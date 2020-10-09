package main

import (
	"fmt"
)

/*
Напишите программу, принимающая на вход число N (N ≥ 4), а затем N целых
чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент
данного среза.

Sample Input:
5
41 -231 24 49 6

Sample Output:
49
*/

func main() {
	var n uint
	fmt.Scan(&n)

	items := make([]int, n)
	for i := range items {
		fmt.Scan(&items[i])
	}
	fmt.Println(items[3])
}
