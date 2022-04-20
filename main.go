package main

import "log"

func getFullName() (string, string) {
	return "Sandrian", "Syafri"
}

func main() {

	firstName, lastName := getFullName()
	log.Println(firstName, lastName)

}
