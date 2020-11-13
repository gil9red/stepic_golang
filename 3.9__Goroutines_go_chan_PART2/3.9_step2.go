package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started!")

	// Ручной wait через канал
	{
		myFunc := func(done chan struct{}) {
			fmt.Print("Hello")
			close(done)
		}

		done := make(chan struct{})
		go myFunc(done)
		<-done
		fmt.Print(" World!")

		fmt.Println()
	}

	// Функция с само-wait
	{
		myFunc := func() <-chan struct{} {
			done := make(chan struct{})
			go func() {
				fmt.Print("Hello")
				close(done)
			}()
			return done
		}
		<-myFunc()
		fmt.Print(" World!")

		fmt.Println()
	}

	fmt.Println("Finish!")
}
