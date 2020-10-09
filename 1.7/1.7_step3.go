package main

import "fmt"

func main() {
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)
	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6

	fmt.Println()

	const (
		Sunday2 = iota
		Monday2
		Tuesday2
		Wednesday2
		Thursday2
		Friday2
		Saturday2
	)
	fmt.Println(Sunday2)   // вывод 0
	fmt.Println(Saturday2) // вывод 6

	const Sunday3 = iota
	const Saturday3 = iota
	fmt.Println(Sunday3, Saturday3) // 0 0

	fmt.Println()

	const (
		N0 = iota
		N1
		_
		N3
		_
		N5
	)
	fmt.Println(N0, N1, N3, N5) // 0 1 3 5

	fmt.Println()

	const (
		A10 = (iota + 1) * 10
		A20
		A30
	)
	fmt.Println(A10, A20, A30) // 10 20 30

	fmt.Println()

	const (
		u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
		v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
		c         = iota * 42 // w == 84  (индекс - 2, поэтому 2 * 42 = 84)
	)
	fmt.Println(u, v, c) // 10 20 30

	fmt.Println()

	const (
		a = iota + 1
		_
		b
		c2
		d = c2 + 2
		t
		i
		i2 = iota + 2
	)
	fmt.Println(a, b, c2, d, t, i, i2) // 1 3 4 6 6 6 9
}
