package main

import "fmt"

func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("error dengan message: ",message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("APLIKASI ERROR")
		// jika menngunakan panic maka aplikasinya akan di stop disini
	}
	fmt.Println("aplikasi berjalan") // dan ini gak bakalan jalan
}

func main(){
	runApp(true)
}