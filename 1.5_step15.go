package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(a / 10 % 10)
}
