package main

import "fmt"

func main(){

	type NoKTP string
	type status bool

	//sekarang untuk mendeklarasikan string bisa menggunakan NoKtp 
	//dan mendeklarasikan boolean menggunakan status 
	var MyKtp NoKTP = "1111112dj333-33"
	var married status = false
	fmt.Println(MyKtp)
	fmt.Println(married)
	fmt.Println(NoKTP("2349491134")) 
}