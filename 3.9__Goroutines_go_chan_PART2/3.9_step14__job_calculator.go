package main

import (
	"fmt"
)

/*
Вам необходимо написать функцию calculator следующего вида:
    func calculator(arguments <-chan int, done <-chan struct{}) <-chan int

В качестве аргумента эта функция получает два канала только для чтения, возвращает
канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о
необходимости завершить работу. Когда сигнал о завершении работы будет получен,
функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.

Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы
этого не сделаете, то превысите предельное время работы.
*/

func main() {
	fmt.Println("Started!")

	calculator := func(arguments <-chan int, done <-chan struct{}) <-chan int {
		result := make(chan int)

		go func() {
			defer close(result)

			var sum int

			for {
				select {
				case value := <-arguments:
					sum += value
				case <-done:
					result <- sum
					return
				}
			}
		}()

		return result
	}

	arguments := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(done)

		for i := 1; i < 10; i++ {
			arguments <- i
		}
	}()

	fmt.Println("Result:", <-calculator(arguments, done))

	fmt.Println("Finish!")
}
