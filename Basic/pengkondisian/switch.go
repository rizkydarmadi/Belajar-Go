package main

import "fmt"

func main(){
	name := "Elia"

	switch name {
	case "ilE":
		fmt.Println("helo Ile")
	case "Eli":
		fmt.Println("hello Eli")
	default:
		fmt.Println("not Found")
	}

	// short statement
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("fizz")
	case false:
		fmt.Println("buzz")
	}

	// tanpa kondisi
	jumlahHuruf := len(name)
	switch{
	case jumlahHuruf > 3:
		fmt.Println("jumlah huruf lebih dari 3")
	case jumlahHuruf < 2:
		fmt.Println("jumlah huruf kurang dari 2")
	default:
		fmt.Println("semua salah")
	}
}