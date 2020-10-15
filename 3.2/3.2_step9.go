package main

import (
	"fmt"
	"strconv"
)

func main() {
	{
		foo := "10"
		bar := "15"
		barInt, err := strconv.Atoi(bar)
		if err != nil {
			panic(err)
		}
		fooInt, err := strconv.Atoi(foo)
		if err != nil {
			panic(err)
		}
		baz := barInt - fooInt
		fmt.Println(baz) // 5

		fmt.Println()
	}

	{
		s := "323str"
		result, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err) // strconv.Atoi: parsing "323str": invalid syntax
		} else {
			fmt.Println(result)
		}

		fmt.Println()
	}

	{
		s := "23.23456"
		// как и в прошлом шаге, здесь 2 параметр - bitSize
		// bitSize - 32 или 64 (32 для float32, 64 для float64)
		// но нужно понимать что метод все равно вернет float64
		result, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)         // 23.23456
		fmt.Printf("%T \n", result) // float64

		// Конкретный пример для разных bitSize:
		s = "1.0000000012345678"
		//  не будем обрабатывать ошибки в примерах, но на практике так не делайте ;)
		result32, _ := strconv.ParseFloat(s, 32)
		result64, _ := strconv.ParseFloat(s, 64)
		fmt.Println("bitSize32:", result32) // вывод 1 (не уместились)
		fmt.Println("bitSize64:", result64) // вывод  1.0000000012345678

		fmt.Println()
	}

	{
		s := "-12345"
		res, err := strconv.ParseInt(s, 10, 64)
		if err != nil { // не забываем проверить ошибку
			panic(err)
		}
		fmt.Println(res) // -12345

		s = "12345"
		res2, err := strconv.ParseUint(s, 10, 64)
		if err != nil { // не забываем проверить ошибку
			panic(err)
		}
		fmt.Println(res2) // 12345

		fmt.Println()
	}

	{
		s := "true"
		res, err := strconv.ParseBool(s)
		if err != nil { // не забываем проверить ошибку
			panic(err)
		}
		fmt.Println(res)      // true
		fmt.Printf("%T", res) // bool
	}
}
