package main

import "fmt"

func getFullName() (firstName,middleName,Lastname string){
	firstName = "son"
	middleName = "heung"
	Lastname = "min"

	return
}

func main(){
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}