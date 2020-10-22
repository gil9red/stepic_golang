package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Умножим на 2:", v*2)
	case string:
		fmt.Println(v + " golang")
	default:
		fmt.Printf("Я не знаю такого типа: '%T'!\n", v)
	}
}

func main() {
	{
		var i interface{} = 12

		if v, ok := i.(int); ok {
			fmt.Println(v + 12) // Суммирование не произойдет, если ok == false
		}

		fmt.Println()
	}

	{
		do(21)
		do("hello")
		do(true)
	}
}
