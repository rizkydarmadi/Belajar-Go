package main

import "fmt"

func getFullName() (string,string,string) {
	return "taylor","swift","test underscore"
}

func main(){
	firstName,lastName,_ := getFullName()
	fmt.Println(firstName,lastName)
}