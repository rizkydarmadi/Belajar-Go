package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist ){
	if blacklist(name){
		fmt.Println("you're Blocked",name)
	} else {
		fmt.Println("welcome",name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main(){

	blacklist := func(name string) bool {
		return name == "admin"
	}

	//anonymous function berarti tidak 

	registerUser("admin", blacklist)
	registerUser("bunga", blacklist)


}