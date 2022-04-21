package main

import "log"

type Person struct {
	Name string
}

func UpdateName(person *Person) {
	person.Name = "Edited"
	log.Println(person.Name)
}

func main() {

	person := Person{"Hafid"}
	UpdateName(&person)

	log.Println(person.Name)

}
