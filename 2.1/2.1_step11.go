package main

import (
	"fmt"
)

/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

func sumInt(args ...int) (int, int) {
	sum := 0
	for _, x := range args {
		sum += x
	}
	return len(args), sum
}

func main() {
	fmt.Println(sumInt(4))
	fmt.Println(sumInt(1, 2, 3, 4))
}
