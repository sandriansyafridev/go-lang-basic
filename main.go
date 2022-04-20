package main

import "log"

func main() {

	// cara pertama
	number := 1
	for number <= 10 {
		log.Println(number)
		number++
	}

	log.Println("================================") //end cara pertama

	//cara kedua
	for i := 1; i <= 10; i++ {
		log.Println(i)
	}

	items := []string{
		"item-1",
		"item-2",
		"item-3",
		"item-4",
	}

	log.Println("================================") //end cara kedua

	for _, item := range items {
		log.Println(item)
	}

}
