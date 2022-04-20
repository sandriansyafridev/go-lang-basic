package main

import "log"

func main() {

	// cara pertama
	var name string
	name = "Sandrian"
	name = "Sandrian Syafri"
	log.Println(name)

	//cara kedua
	var address string = "Jambi"
	log.Println(address)

	// cara ketiga
	telepon := "08xxx"
	log.Println(telepon)

	//cara ke empat
	var (
		email    = "sandriansyafri.dev@gmail.com"
		password = "password"
	)
	log.Println(email)
	log.Println(password)

}
