package main

import "fmt"

type Animal struct {
	Name string
	Age  uint
}

func (a *Animal) String() string {
	return fmt.Sprintf("<Animal Name=%v Age=%d>", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
	// {Gopher 2}

	fmt.Println(&a)
	// <Animal name=Gopher age=2>

	fmt.Println()

	aP := &Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(aP)
	// <Animal name=Gopher age=2>

	fmt.Println(*aP)
	// {Gopher 2}
}
