package main

import "fmt"

func main() {
	var x int64 = 57
	var y = float64(x)
	fmt.Print(x, y) // 57 57

	fmt.Println()

	var f float64 = 56.231
	var i = int(f)
	fmt.Println(f, i) // 56.231 56

	fmt.Println()

	fmt.Println(5/2, 5.0/2) // 2 2.5
}
