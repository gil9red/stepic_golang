package main

import (
	"fmt"
)

func externalFunction() func() {
	text := "TEXT"

	return func() {
		fmt.Println(text)
	}
}

func externalFunctionV2(text string) func() {
	return func() {
		fmt.Println(text)
	}
}

func main() {
	{
		fn := externalFunction()
		fn()
		// TEXT

		fnHello := externalFunctionV2("Hello World!")
		fnHello()
		fnHello()
		// Hello World!
		// Hello World!

		fmt.Println()
	}

	{
		fn := func() func() int {
			count := 0
			return func() int {
				count++
				return count
			}
		}()

		for i := 1; i <= 5; i++ {
			fmt.Print(fn())
		}
		fmt.Println()
		// 12345

		fmt.Println()
	}

	{
		counterFabric := func(count int) func() int {
			return func() int {
				count++
				return count
			}
		}

		fn1 := counterFabric(0)
		fmt.Println(fn1(), fn1(), fn1(), fn1(), fn1())
		// 1 2 3 4 5

		fn2 := counterFabric(10)
		fmt.Println(fn2(), fn2(), fn2(), fn2(), fn2())
		// 11 12 13 14 15

		fmt.Println()
	}

	{
		fn := func() func(int) int {
			count := 0
			return func(i int) int {
				count++
				return count * i
			}
		}()

		for i := 1; i <= 5; i++ {
			fmt.Print(fn(i), " ")
		}
		// 1 4 9 16 25

		fmt.Println()
	}
}
