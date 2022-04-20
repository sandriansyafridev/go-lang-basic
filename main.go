package main

import "log"

func filterName(value string) string {
	switch {
	case value != "anjing":
		return value
	default:
		return "..."
	}
}

func displayName(value string, filterName func(string) string) string {
	return filterName(value)
}

func main() {

	result1 := displayName("anjing", filterName)
	result2 := displayName("Sandrian Syafri", filterName)
	log.Println(result1)
	log.Println(result2)

}
