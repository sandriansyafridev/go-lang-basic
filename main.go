package main

import "log"

func main() {

	// cara satu - inline declaration
	const APP = "Go APP"
	const VERSION = "1.0.0"
	log.Println(APP, VERSION)

	//cara dua - multiple declaration
	const (
		FIRST_NAME = "Sandrian"
		LAST_NAME  = "Syafri"
	)
	log.Println(FIRST_NAME, LAST_NAME)

}
