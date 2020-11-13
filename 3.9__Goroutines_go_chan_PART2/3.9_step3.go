package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started!")

	work := func() {
		fmt.Println("Hello World!")
	}

	myFunc := func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			work()
			close(done)
		}()
		return done
	}
	<-myFunc()

	fmt.Println("Finish!")
}
