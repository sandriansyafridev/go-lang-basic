package main

import "log"

func main() {

	// cara 1
	var myStringArr [2]string
	myStringArr[0] = "item 1"
	myStringArr[1] = "item 2"
	myStringArr[0] = "item 1 edited" //update value on index 0

	log.Println("my string array = ", myStringArr)
	log.Println(len(myStringArr)) //length array

	// cara 2
	var myIntArr = [2]int{1, 2}
	log.Println("my int array = ", myIntArr)

	// cara 3
	myArrWithTypeInterface := [3]interface{}{"Hello World", 1, true}
	log.Println("my array = ", myArrWithTypeInterface)

	//cara 4
	myArrWithUndefinedLenght := [...]int{
		1, 2, 3, 4, 5,
	}

	log.Println(myArrWithUndefinedLenght)

}
