package main

import "log"

func DisplayAnything() interface{} {
	return 10
}

func main() {

	result := DisplayAnything()

	switch value := result.(type) {
	case string:
		log.Println(value)
	default:
		log.Println("not string")
	}
}
