package main

import (
	"fmt"
	"strings"
)

/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в
строке X. Если подстроки S нет в строке X - вывести -1

Sample Input:
awesome
es

Sample Output:
2
*/

func main() {
	var text1, text2 string
	_, _ = fmt.Scan(&text1, &text2)

	fmt.Println(strings.Index(text1, text2))
}
