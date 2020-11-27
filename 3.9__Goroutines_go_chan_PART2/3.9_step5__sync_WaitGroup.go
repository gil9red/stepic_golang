package main

import (
	"fmt"
	"sync"
)

/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах
вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.

Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
*/

func main() {
	fmt.Println("Started!")

	work := func() {
		// Будет определена в stepic
	}

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")

	fmt.Println("Finish!")
}
