package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("my error")
	fmt.Println(err)
	// my error

	err2 := fmt.Errorf("my error: %s", "<error>")
	fmt.Println(err2)
	// my error: <error>
}
