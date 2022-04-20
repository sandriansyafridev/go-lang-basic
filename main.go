package main

import "log"

func displayName(value string, filter func(value string) string) string {
	return filter(value)
}

func main() {
	result := displayName("Sandrian", func(value string) string {
		switch {
		case value != "anjing":
			return value
		default:
			return "..."
		}
	})

	log.Println(result)

}
