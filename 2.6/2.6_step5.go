package main

import (
	"fmt"
)

func dividePanic(a, b float64) float64 {
	if b == 0 {
		panic("division by zero!")
	}
	return a / b
}

func main() {
	fmt.Println(dividePanic(15, 5))
	fmt.Println(dividePanic(4, 0))
	fmt.Println("Program has been finished")
}
