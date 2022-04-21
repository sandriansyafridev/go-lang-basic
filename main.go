package main

import (
	"errors"
	"log"
)

func Devide(pembilang int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("cannot devide")
	}

	result := pembilang / pembagi
	return result, nil
}

func main() {

	result, err := Devide(10, 1)
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println(result)

}
