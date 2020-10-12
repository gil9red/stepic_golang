package main

import "fmt"

func main() {
	rs := []rune("Это срез рун")

	// Итерируясь мы будем заменять символ 'р' на '*'
	for i := range rs {
		switch rs[i] {
		case 'р':
			rs[i] = '*'
		case ' ':
			rs[i] = '_'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))
}
