package main

import "fmt"

/*
Требуется написать программу, при выполнении которой с клавиатуры считываются
два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B включительно.

Sample Input:
1 5

Sample Output:
15
*/

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	n := 0
	for i := a; i <= b; i++ {
		n += i
	}
	fmt.Println(n)
}
