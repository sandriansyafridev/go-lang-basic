package main

import "log"

func displayHelloWorld() {
	log.Println("Hello World!")
}

func main() {

	displayResult := displayHelloWorld
	displayResult()

}
