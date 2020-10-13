package main

import (
	"fmt"
)

func finish() {
	fmt.Println("Program has been finished")
}

func myFunc(a interface{}) {
	fmt.Println("myFunc", a)
}

func main() {
	defer finish()

	a := 3
	defer myFunc(a)

	a = 2
	defer myFunc(a)

	defer func() {
		for i := 0; i < 3; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println("!")
	}()

	defer fmt.Println("Program has been started")

	fmt.Println("Program is working")

	myFunc(1)

	panic("panic! panic! panic!")
}
