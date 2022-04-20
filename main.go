package main

import "log"

func main() {

	var number1 int32 = 100
	var number2 int64 = int64(number1)
	log.Println(number1)
	log.Println(number2)

	var name string = "Sandrian Syafri"
	var index1 string = string(name[0])
	log.Println(name)
	log.Println(index1)

}
