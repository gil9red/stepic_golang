package main

import (
	"fmt"
)

/*
Вам необходимо написать функцию calculator следующего вида:
    func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int

Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
 * в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал
   вы должны отправить квадрат аргумента.
 * в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал
   вы должны отправить результат умножения аргумента на 3.
 * в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.

Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит
всего одно значение в один из каналов - получили значение, обработали его, завершили работу.

После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого
не сделаете, то превысите предельное время работы.
*/

func main() {
	fmt.Println("Started!")

	calculator := func(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
		result := make(chan int)

		go func() {
			defer close(result)

			select {
			case value := <-firstChan:
				result <- value * value
			case value := <-secondChan:
				result <- value * 3
			case <-stopChan:
				return
			}
		}()

		return result
	}

	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		defer close(stopChan)

		firstChan <- 10
		firstChan <- 100
		secondChan <- 10
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Result:", <-calculator(firstChan, secondChan, stopChan))
	}
	// Result: 100
	// Result: 10000
	// Result: 30
	// Result: 0

	fmt.Println("Finish!")
}
