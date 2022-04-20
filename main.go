package main

import "log"

func getFullName() (firstName string, lastName string) {

	firstName = "Sandrian"
	lastName = "Syafri"

	return
}

func main() {

	firstName, lastName := getFullName()
	log.Println(firstName, lastName)

}
