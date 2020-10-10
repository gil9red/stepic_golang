package main

import (
	"fmt"
)

/*
Напишите функцию f(), которая будет принимать строку text и выводить ее
(печатать на экране).
От вас требуется дописать только эту функцию, считайте что функция main()
уже объявлена, считывать с консоли ничего не нужно!
Пакет "fmt" уже импортирован!

Sample Input:
Hello!

Sample Output:
Hello!
*/

func f(text string) {
	fmt.Println(text)
}

func main() {
	text := "Hello World"
	f(text)
}
