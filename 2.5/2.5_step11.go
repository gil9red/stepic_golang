package main

import (
	"fmt"
	"strings"
)

/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза
и вывести получившуюся строку

Sample Input:
zaabcbd

Sample Output:
zcd
*/

func main() {
	var text string
	_, _ = fmt.Scan(&text)

	var b strings.Builder
	for _, r := range text {
		if strings.Count(text, string(r)) == 1 {
			b.WriteRune(r)
		}
	}

	fmt.Println(b.String())
}
