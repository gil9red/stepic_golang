package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	fmt.Println(a + b)
}
