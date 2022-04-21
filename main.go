package main

import "log"

type HasName interface {
	GetName()
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() {
	log.Println(person.Name)
}

func (animal Animal) GetName() {
	log.Println(animal.Name)
}

func DisplayName(f HasName) {
	f.GetName()
}

func main() {

	person := Person{"Sandrian"}
	animal := Animal{"Cat"}

	DisplayName(person)
	DisplayName(animal)

}
