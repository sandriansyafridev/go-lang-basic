package main

import "log"

type Person struct {
	Firstname string
	Lastname  string
	Address   string
	Age       int
}

func main() {

	//cara satu
	person1 := Person{
		Firstname: "Sandrian",
		Lastname:  "Syafri",
		Address:   "Muara Bungo",
	}
	log.Println(person1.Firstname, person1.Lastname, person1.Address)

	//cara dua
	var person2 Person
	person2.Firstname = "Hafid"
	person2.Lastname = "Dwi"
	person2.Address = "Padang"
	log.Println(person2.Firstname, person2.Lastname, person2.Address)

	//cara tiga
	person3 := Person{"Yogi", "Okamura", "Padang", 28}
	log.Println(person3.Firstname, person3.Lastname, person3.Address, person3.Age)

}
