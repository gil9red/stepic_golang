package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Itoa
	{
		text := strconv.Itoa(2020)        // int -> string
		fmt.Printf("%T %v\n", text, text) // string 2020

		fmt.Println()
	}

	// FormatInt
	{
		var n int64 = 0xB                     // 'B' в шестнадцатеричной это 11 в десятичной системе
		fmt.Println(strconv.FormatInt(n, 2))  // 1011
		fmt.Println(strconv.FormatInt(n, 10)) // 11
		fmt.Println(strconv.FormatInt(n, 16)) // b

		fmt.Println()
	}

	// FormatUint
	{
		var n uint64 = 93101
		res := strconv.FormatUint(n, 10)
		fmt.Printf("%T %v\n", res, res) // string 93101

		fmt.Println()
	}

	// FormatFloat
	{
		var a float64 = 1.0123456789

		// 1 параметр - число для конвертации
		// fmt - форматирование
		// prec - точность (кол-во знаков после запятой)
		// bitSize - 32 или 64 (32 для float32, 64 для float64)
		fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

		// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
		fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

		// Возможные форматы fmt:
		// 'f' (-ddd.dddd, no exponent),
		// 'b' (-ddddp±ddd, a binary exponent),
		// 'e' (-d.dddde±dd, a decimal exponent),
		// 'E' (-d.ddddE±dd, a decimal exponent),
		// 'g' ('e' for large exponents, 'f' otherwise),
		// 'G' ('E' for large exponents, 'f' otherwise),
		// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
		// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
		var b float64 = 2222 * 1023 * 245 * 2 * 52
		fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10

		fmt.Println()
	}

	// Sprint and Sprintf
	{
		fmt.Println(fmt.Sprint(20.19)) // Краткая форма

		a := 20.20
		fmt.Println(fmt.Sprintf("%f", a)) // Полная форма

		fmt.Println()
	}

	// FormatBool
	{
		a := true
		res := strconv.FormatBool(a)
		fmt.Printf("%T %v", res, res) // string true

		fmt.Println()
	}
}
