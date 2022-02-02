package main

import "fmt"

func main(){

	counter := 0
	name := "auba"

	increment := func(){
		name = "vlahovic"
		fmt.Println("increment")
		counter++
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)

}