package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi Selesai")
	meessage_error := recover() //! recover digunakan untuk error handiling biar erronya ketangkep dan program terus berlanjut 
	//! harus ditaruh di luar fungsi yang bakalan error 
	//! karena defer bakalan dijalanin terakhir maka recover lebih baik ditaro di defer
	if meessage_error != nil {
		fmt.Println("terjadi error", meessage_error)
	}

	fmt.Println("aplikasi Berjalan")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
}

func main(){
	runApp(false)
}