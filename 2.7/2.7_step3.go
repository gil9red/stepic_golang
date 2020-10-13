package main

import (
	"fmt"
)

/*
Дана строка, содержащая только английские буквы (большие и маленькие).
Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после
последней символ ‘*’ добавлять не нужно).

Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает
1000 знаков.

Выходные данные
Вывести строку, которая получится после добавления символов '*'.

Sample Input:
LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO

Sample Output:
L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O

Sample Input:
LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO

Sample Output:
L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
*/

func main() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid input values")
	}
	bs := []rune(text)

	fmt.Print(string(bs[0]))
	for i := 1; i < len(bs); i++ {
		fmt.Print("*", string(bs[i]))
	}
}
