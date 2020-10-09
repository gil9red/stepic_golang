package main

import (
	"fmt"
)

/*
Заданы три числа - a,b,c (a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

Sample Input:
6 8 10

Sample Output:
Прямоугольный
*/

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
