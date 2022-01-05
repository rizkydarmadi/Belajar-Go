package main

import "fmt"

func main(){
	//method 1
	var name string

	name = "Jennie"
	fmt.Println(name)
	name = "Jennie Kim"
	fmt.Println(name)

	//method 2
	var name2 = "rose"
	fmt.Println(name2)
	name2 = "roseanne park"
	fmt.Println(name2)

	//method 3
	name3 := "jiso"
	fmt.Println(name3)
	name3 = "jisoya"
	fmt.Println(name3)

	//method 4
	var (
		firstName = "lalisa"
		lastname = "manoban"
	)
	fmt.Println(firstName)
	fmt.Println(lastname)
}