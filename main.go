package main

import "log"

func main() {

	//create map value
	mapString := map[string]string{
		"first_name": "Sandrian",
		"last_name":  "Syafri",
		"address":    "Jambi",
	}

	mapString["first_name"] = "edited" // update map value
	log.Println(mapString, "len map before delete =", len(mapString))

	delete(mapString, "address")
	log.Println(mapString)
	log.Println(mapString, "len map after delete =", len(mapString))
}
