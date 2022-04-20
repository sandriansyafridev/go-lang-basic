package main

import "log"

func main() {

	myItem := [...]string{
		"item 1",
		"item 2",
		"item 3",
		"item 4",
		"item 5",
		"item 6",
		"item 7",
		"item 8",
		"item 9",
		"item 10",
	}

	// cara 1 - array[low:high]
	// mengambil nilai dari index low hingga sebelum index high
	firstWaySlice := myItem[0:4]
	log.Println(firstWaySlice)

	//cara 2 - array[low:]
	//mengambil nilai dari index low hingga akhir index high
	secondWaySlice := myItem[0:]
	log.Println(secondWaySlice)

	//cara 3 - array[:high]
	//mengambil semua nilai hingga index sebelum high
	thirdWaySlice := myItem[:5]
	log.Println(thirdWaySlice)

	//cara ke 4 - array[:]
	//mengambil semua nilai hingga index high
	fourthWaySlice := myItem[:]
	log.Println(fourthWaySlice)

	// cara ke 5 - membuat slice baru
	newSlice1 := []string{
		"list 1",
		"list 2",
		"list 3",
	}
	newSlice1 = append(newSlice1, "list 4")
	newSlice1 = append(newSlice1, "list 5")
	newSlice1 = append(newSlice1, "list 6")
	log.Println(newSlice1)

}
