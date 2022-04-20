package main

import "log"

func main() {

	// break statment
	for i := 1; i <= 10; i++ {
		log.Println(i)
		if i == 5 {
			break
		}
	}

	log.Println("=======================") // end break statment

	//continue stetment
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		log.Println(i)
	}

}
