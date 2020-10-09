package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	a *= 2
	a += 100
	fmt.Println(a)
}
