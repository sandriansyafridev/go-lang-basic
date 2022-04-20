package main

import "log"

func main() {

	x := 10
	y := 11

	log.Println("===========")
	log.Println("x =", x)
	log.Println("y =", y)
	log.Println("===========")
	log.Println("x == y =", x == y)
	log.Println("x != y =", x != y)
	log.Println("x > y =", x > y)
	log.Println("x >= y =", x >= y)
	log.Println("x < y =", x < y)
	log.Println("x <= y =", x <= y)

}
