package main

import "fmt"

func main(){

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulanan ke = ", counter)
	// 	counter++
	// }

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke = ",counter)
	}

	slice := []string{"lisa","jiso","rose","rose"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("ini adalah = ",slice[i])
	}

	for index, name := range slice {
		fmt.Println("index", index , "=", name)
	}

}