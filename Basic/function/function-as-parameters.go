package main

import "fmt"

type Filter func(string) string

func sayHelloWithfilter(name string,filter Filter){
	fmt.Println("hello", filter(name))
}


func spamFilter(name string) string{
	if name == "anjing" {
		return "..."
	}else{
		return name
	}
}

func main(){
	// sayHelloWithfilter("jack",spamFilter)
	sayHelloWithfilter("anjing",spamFilter)
}
