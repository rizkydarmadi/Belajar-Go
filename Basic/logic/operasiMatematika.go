package main

import "fmt"

func main(){
	// operasi matematika sederhana
	var a = 10
	var b = 10
	var c = a + b

	const d = 10 + 16
	
	fmt.Println(a,"+",b,"=",c)

	//augmented assigment
	var x = 0
	var y = x + 13
			x += 13
			
	fmt.Println(x,y)
	//unary operator
	var j = 5
	j++ // j = j + 1

	fmt.Println(j)
	


	
}