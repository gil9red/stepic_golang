package main

import (
	"errors"
	"fmt"
)

/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает
не только результат деления, но и информацию об ошибке. В случае какой-либо
ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
Функция divide(a int, b int) (int, error) получает на вход два числа которые
нужно поделить и возвращает результат (int) и ошибку (error).
Не забудьте считать два числа которые необходимо поделить (передать в функцию)!

Sample Input:
10 5

Sample Output:
2
*/

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}

	return a / b, nil
}

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	c, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	fmt.Println(c)

	//fmt.Println(divide(4, 2))
	//fmt.Println(divide(4, 0))
}
