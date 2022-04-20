package main

import (
	"fmt"
	"log"
)

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {

	displaySayHello := sayHello("Sandrian Syafri")
	log.Println(displaySayHello)

}
