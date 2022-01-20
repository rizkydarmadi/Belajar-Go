package main

import "fmt"

func main(){
	name := "willian"

	if name == "julian"{
		fmt.Println("this is Julian")
	}else if name == "willian"{
		fmt.Println("this is willy")
	}else{
		fmt.Println("name not found")
	}

	if lenght := len(name);lenght > 5{
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("nama sudah benar")
	}
}

