package main

import "log"

type Person struct {
	Name string
}

func (person *Person) UpdateName(value string) {
	person.Name = value
}

func main() {

	person := Person{"Hafid"}

	log.Println(person)
	person.UpdateName("Hafid edited!")
	log.Println(person)

}
