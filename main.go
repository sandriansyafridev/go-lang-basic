package main

import "log"

func main() {

	name := "Sandrian Syafri"

	// cara satu - if else if
	if name == "Sandrian Syafri" {
		log.Println("Hello", name)
	} else if name == "Hafid" {
		log.Println("Hi", name)
	} else {
		log.Println("No found!")
	}

	// cara dua - if with short statment
	if len := len(name); len > 5 {
		log.Println("Besar dari 5")
	} else {
		log.Println("Kecil dari 5")
	}

}
