package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Started!")

	// Пример ограничения времени жизни выполнения кода. Последовательное выполнение
	{
		c := make(chan int)
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
				time.Sleep(time.Millisecond * 300)
			}
		}()

		timeout := time.After(2 * time.Second)

	LOOP:
		for {
			//for i := 0; i < 5; i++ {
			select { // Оператор select
			case gopherID := <-c: // Ждет, когда проснется гофер
				fmt.Println("gopher", gopherID, "has finished sleeping")

			case <-timeout: // Ждет окончания времени
				fmt.Println("my patience ran out")
				break LOOP
				//return // Сдается и возвращается
			}
		}

		fmt.Println()
	}

	// Еще пример ограничения времени жизни выполнения кода + рандом задержки. Параллельное выполнение
	{
		c := make(chan int)

		sleepyGopher := func(id int, c chan int) {
			time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
			c <- id
		}

		for i := 0; i < 100; i++ {
			go sleepyGopher(i, c)
		}

		timeout := time.After(2 * time.Second)

	LOOP2:
		for {
			//for i := 0; i < 5; i++ {
			select { // Оператор select
			case gopherID := <-c: // Ждет, когда проснется гофер
				fmt.Println("gopher", gopherID, "has finished sleeping")

			case <-timeout: // Ждет окончания времени
				fmt.Println("my patience ran out")
				break LOOP2
				//return // Сдается и возвращается
			}
		}

		fmt.Println()
	}

	// select - case - default
	{
		tick1 := time.After(time.Second)
		tick2 := time.After(time.Second * 2)
		select {
		case <-tick1:
			fmt.Println("Получено значение из первого канала")
		case <-tick2:
			fmt.Println("Получено значение из второго канала")
			// Блок default выполнится раньше блока case - 1 секунда слишком много для Go
		default:
			fmt.Println("Действие по умолчанию")
		}
		// Действие по умолчанию
	}

	fmt.Println("Finish!")
}
