package main

import "fmt"

func logging(){
	fmt.Println("selesai menjalankan program")
}

func runApp(value int){
	defer logging()
	fmt.Println("run app")
	result := 10 / value
	fmt.Println(result)
}

func main(){
	runApp(0)
}