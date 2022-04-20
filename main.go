package main

import "log"

func main() {

	nilaiUts := 70
	nilaiUas := 80

	if nilaiUts > 75 && nilaiUas > 75 {
		log.Println("PASS EXAM")
	} else {
		log.Println("NOT PASS EXAM")
	}

}
