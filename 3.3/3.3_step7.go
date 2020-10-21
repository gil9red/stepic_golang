package main

import (
	"fmt"
	"strconv"
)

/*
Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида
func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать
по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число
того же типа в котором отсутствуют нечетные цифры или цифра 0, если же получившееся
число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же
порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:
727178

Sample Output:
28
*/

func main() {
	fn := func(num uint) uint {
		numStr := strconv.FormatUint(uint64(num), 10)
		rs := make([]rune, 0)

		for _, r := range numStr {
			if r%2 == 0 {
				rs = append(rs, r)
			}
		}
		numStr = string(rs)

		x, _ := strconv.Atoi(numStr)
		if x == 0 {
			x = 100
		}

		return uint(x)
	}

	fmt.Println(fn(727178))
	// 28
}
