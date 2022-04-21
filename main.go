package main

import "log"

type Person struct {
	Name string
}

func (person Person) SayHello() {
	log.Println(person.Name)
}

func main() {

	person := Person{"Sandrian"}
	person.SayHello()

}
