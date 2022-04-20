package main

import "log"

func initApp() {
	log.Println("START RUNNING APPLICATION")

}

func runApp(err bool) {
	if err {
		panic("SOMETHING WRONG")
	}

	log.Println("RUNNING APPLICATION")
	log.Println("RUNNING APPLICATION")
	log.Println("RUNNING APPLICATION")
	log.Println("RUNNING APPLICATION")

}

func endApp() {
	errorMessage := recover()

	log.Println("FINISH RUNNING APPLICATION")
	log.Println(errorMessage)
}

func main() {

	defer endApp() // dieksekusi setelah semua proses selesai.

	initApp()
	runApp(true)

}
