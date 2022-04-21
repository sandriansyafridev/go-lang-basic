package main

import "log"

func DisplayName(name string) interface{} {
	if name == "" {
		return nil
	}

	return name
}

func main() {

	if result := DisplayName(""); result != nil {
		log.Println(result)
	}

	log.Println("no value")

}
