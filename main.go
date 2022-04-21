package main

import "log"

type Person struct {
	Name string
	Age  int
}

func main() {

	person1 := Person{"Sandrian", 23}
	person2 := &person1

	*person2 = Person{"Hafid", 24}

	log.Println(person1)
	log.Println(person2)

}
