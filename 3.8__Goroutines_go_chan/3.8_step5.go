package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Started!")

	{
		c := make(chan string)

		f := func(c chan string) {
			c <- "1"
			c <- "2"
			c <- "3"
			close(c)
		}
		go f(c)

		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}

		fmt.Println()
	}

	{
		const N = 5

		worker := func(wg *sync.WaitGroup, c chan string, name string, timeout time.Duration) {
			defer wg.Done()
			wg.Add(1)

			for i := 1; i <= N; i++ {
				c <- fmt.Sprintf("[%v] %v", name, i)
				time.Sleep(timeout)
			}
		}

		monitorWorker := func(wg *sync.WaitGroup, cs chan string) {
			wg.Wait()
			close(cs)
		}

		wg := &sync.WaitGroup{}
		c := make(chan string, 10)

		go worker(wg, c, "a", time.Millisecond*100)
		go worker(wg, c, "b", time.Millisecond*300)
		go worker(wg, c, "c", time.Millisecond*50)

		go monitorWorker(wg, c)

		for x := range c {
			fmt.Println(x)
		}
	}

	fmt.Println("Finish!")
}
