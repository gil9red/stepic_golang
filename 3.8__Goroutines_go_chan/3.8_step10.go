package main

import (
	"fmt"
)

/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения
на следующий этап конвейера только если оно отличается от того, что пришло ранее.

Ваша фукнция должна принимать два канала - inputStream и outputStream, в первый вы будете
получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream
должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
*/

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	prevValue := <-inputStream
	outputStream <- prevValue

	for value := range inputStream {
		if prevValue == value {
			continue
		}

		outputStream <- value
		prevValue = value
	}
}

func main() {
	fmt.Println("Started!")

	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
	// 123456

	fmt.Println("\nFinish!")
}
