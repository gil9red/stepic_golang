package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Started!")

	{
		t := time.NewTimer(time.Second) // создаем новый таймер, который сработает через 1 секунду
		go func() {
			<-t.C // C - канал, который должен вернуть значение через заданное время
		}()
		t.Stop() // но мы можем остановить таймер и раньше установленного времени

		t.Reset(time.Second * 2) // пока таймер не сработал, мы можем сбросить его, установив новый срок выполнения
		<-t.C
		fmt.Println("Прошло 2 секунды")
		fmt.Println()
	}

	{
		work := func() <-chan struct{} {
			done := make(chan struct{}) // канал для синхронизации горутин

			go func() {
				defer close(done) // синхронизирующий канал будет закрыт, когда функция завершит свою работу

				stop := time.NewTimer(time.Second)
				tick := time.NewTicker(time.Millisecond * 200)
				defer tick.Stop() // освободим ресурсы, при завершении работы функции

				for {
					select {
					case <-stop.C:
						// stop - Timer, который через 1 секунду даст сигнал завершить работу
						return
					case <-tick.C:
						// tick - Ticker, посылающий сигнал выполнить работу каждый 200 миллисекунд
						fmt.Println("Тик-так")
					}
				}
			}()

			return done
		}
		<-work()

		// Тик-так
		// Тик-так
		// Тик-так
		// Тик-так
		// Тик-так
		fmt.Println()
	}

	{
		worker := func(id int, limit <-chan time.Time, wg *sync.WaitGroup) {
			defer wg.Done()
			<-limit
			fmt.Printf("Worker %d выполнил работу\n", id)
		}

		tick := time.NewTicker(time.Second)
		defer tick.Stop()

		wg := new(sync.WaitGroup)

		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go worker(i, tick.C, wg)
		}

		wg.Wait()

		// Worker 1 выполнил работу
		// Worker 3 выполнил работу
		// Worker 4 выполнил работу
		// Worker 5 выполнил работу
		// Worker 2 выполнил работу
	}

	fmt.Println("Finish!")
}
