package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a)
	fmt.Println(&a)

	var b *int
	b = &a
	fmt.Println(*b)
}