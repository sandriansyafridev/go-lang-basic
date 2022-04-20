package main

import "log"

func main() {

	type NOBP string
	type IS_STUDENT bool

	var noBp NOBP = "1610140"
	log.Println(noBp)

	var isStudent IS_STUDENT = false
	log.Println(isStudent)

}
