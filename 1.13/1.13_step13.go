package main

import (
	"fmt"
	"log"
)

/*
По данному числу N распечатайте все целые значения степени двойки, не
превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
50

Sample Output:
1 2 4 8 16 32
*/

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < n; i++ {
		x := 1 << i
		if x > n {
			break
		}
		fmt.Print(x, " ")
	}
}
