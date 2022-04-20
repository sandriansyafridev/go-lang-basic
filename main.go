package main

import (
	"log"
)

func main() {

	name := "Sandrian Syafri"

	// cara satu
	switch name {
	case "Sandrian Syafri":
		{
			log.Println("Hello", name)
			log.Println("Hello", name)
			log.Println("Hello", name)

		}
	case "Hafid":
		{
			log.Println("Hi ", name)
		}
	default:
		log.Println("Not Found")
	}

	log.Println("====================")

	//cara dua - with short statment
	switch len := len(name); len > 5 {
	case len > 5:
		{
			log.Println("Besar dari 5")
		}
	default:
		log.Println("Kecil dari lima")
	}

	log.Println("====================")

	// cara tiga - with no expression
	switch {
	case name == "Sandrian Syafri":
		{
			log.Println("Hello", name)
		}

	case name == "Hafid":
		{
			log.Println("Hi", name)
		}

	default:
		log.Println("Not Found!")
	}

}
