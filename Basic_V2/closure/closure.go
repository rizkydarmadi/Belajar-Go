package main

import "fmt"

//todo closure adalah kemampuan memanipulasi data
//todo didalam scope { }
func main(){
	counter := 0

	increment := func ()  {
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
}