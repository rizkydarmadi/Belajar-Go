package main

import "fmt"


func sayHello(name string) string {
	return "hello" + name
}

func main(){
	fmt.Println(sayHello(" jennie"))
}