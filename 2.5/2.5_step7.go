package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong

Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы можете
считать полностью строку за раз с помощью bufio:
text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

Sample Input:
Быть или не быть.

Sample Output:
Right
*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	bs := []rune(text)

	if unicode.IsUpper(bs[0]) && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
