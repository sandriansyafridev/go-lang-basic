package main

import "log"

func sumNumbers(numbers ...int) (result int) {

	for _, number := range numbers {
		result += number
	}

	return result

}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	result := sumNumbers(numbers...)
	log.Println(result)

}
