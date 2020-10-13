package main

import (
	"fmt"
)

func main() {
	var m map[int]int
	fmt.Printf("m is nil -> %t", m == nil) // Warning! m is nil!

	fmt.Println()

	// с помощью встроенной функции make:
	m1 := make(map[int]int)

	m2 := map[int]int{}

	// с помощью использования литерала отображения:
	m3 := map[int]int{
		// Пары ключ:значение указываются при необходимости
		12: 2,
		1:  5,
	}

	fmt.Println(m1) // map[]
	fmt.Println(m2) // map[]
	fmt.Println(m3) // map[1:5 12:2]

	fmt.Println()

	m1[1], m2[1], m3[1] = 123, 123, 123
	fmt.Println(m1) // map[1:123]
	fmt.Println(m2) // map[1:123]
	fmt.Println(m3) // [1:123 12:2]
}
