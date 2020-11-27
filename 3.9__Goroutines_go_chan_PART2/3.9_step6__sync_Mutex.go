package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Started!")

	const N = 1000

	// NOTE: WRONG. Result x will not 1000
	{
		var x int
		wg := new(sync.WaitGroup)

		for i := 0; i < N; i++ {
			// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				x++
			}(wg)
		}

		wg.Wait()

		// По идее значение счетчика должно быть 1000, но крайне вероятно,
		// что этого не произойдет
		fmt.Println(x)
	}

	// That's right.
	{
		var x int
		wg := new(sync.WaitGroup)
		mu := new(sync.Mutex)

		for i := 0; i < N; i++ {
			// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
			wg.Add(1)
			go func(wg *sync.WaitGroup, mu *sync.Mutex) {
				defer wg.Done()
				mu.Lock()
				x++
				mu.Unlock()
			}(wg, mu)
		}

		wg.Wait()
		fmt.Println(x)
	}

	fmt.Println("Finish!")
}
