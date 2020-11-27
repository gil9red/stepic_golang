package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Started!")

	work := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)      // Увеличиваем счетчик горутин в группе
		go work(i, wg) // Вызываем функцию work в отдельной горутине
	}

	wg.Wait() // Ожидание завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")

	fmt.Println("Finish!")
}
