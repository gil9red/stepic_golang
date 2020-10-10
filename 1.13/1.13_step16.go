package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272
*/

func main() {
	var n, d int
	_, err := fmt.Scan(&n, &d)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(strings.Replace(strconv.Itoa(n), strconv.Itoa(d), "", -1))
}
