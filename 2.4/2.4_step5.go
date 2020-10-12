package main

import "fmt"

/*
Встраиваемые типы
*/

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	// Person Person  // Composition
	Person // Embedding (analogue of inheritance)
	Model  string
}

func main() {
	android := &Android{
		Person: Person{
			Name: "Vasya",
		},
		Model: "001",
	}
	fmt.Printf("%T, %v\n", android, android)
	fmt.Println(android.Name, android.Person.Name)
	android.Talk()
}
