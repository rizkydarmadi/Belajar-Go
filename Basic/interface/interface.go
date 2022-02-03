package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(hasname HasName) {
	fmt.Println("hello", hasname.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "eko"

	SayHello(eko)

	var animal Animal
	animal.Name = "ciro"

	SayHello(animal)

}
