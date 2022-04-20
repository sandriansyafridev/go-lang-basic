package main

import "log"

func main() {
	x := 5
	y := 10

	//add
	handlePlus := x + y
	log.Println("x + y = ", handlePlus)

	//min
	handleMinus := y - x
	log.Println("x - y = ", handleMinus)

	//muliply
	handleMutiply := x * y
	log.Println("x * y = ", handleMutiply)

	//devide
	handleDevide := y / x
	log.Println("x / y = ", handleDevide)

	//augmented assignment
	a := 5
	a += 10
	// a = 5 + 10
	// a = 15
	log.Println("a += 10 = ", a)

	//increment // unary operation
	b := 1
	b++
	log.Println("b++ = ", b)

	c := 10
	c--
	log.Println("c-- = ", c)

	d := +10 //positive
	log.Println("d = ", d)
	e := -10 //negative
	log.Println("e = ", e)

}
