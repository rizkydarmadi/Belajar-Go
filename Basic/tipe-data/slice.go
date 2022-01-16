package main

import "fmt"

func main(){
	var months = [...]string{
		"january", 
		"febuary",
		"march",
		"april",
		"may",
		"juni",
		"july",
		"august",
		"september",
		"october",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[4] = "rajab"
	// fmt.Println(slice1)

	// fmt.Println(months)

	days := [...]string{"senin", "selasa","rabu","kamis","jum'at","sabtu","minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "saturday"
	daysSlice1[1] = "sunday"

	fmt.Println(days)

	daysSlice2 := append(daysSlice1,"monday")
	daysSlice2[0] = "sabtu"
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice := make([]string, 2,5)

	newSlice[0] = "smith"
	newSlice[1] = "rowe"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,4,5}
	inislice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(inislice)




}