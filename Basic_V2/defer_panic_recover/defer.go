package main

import "fmt"

func logging(){
	fmt.Println("selesai menjalankan program")
}

//! defer digunakan untuk memanggil fungsi sesudah fungsi utamanya
//! kelar, walaupun terdapat error dia tetap dijalankan

func runApp(value int){
	defer logging()
	fmt.Println("run app")
	result := 10 / value
	fmt.Println(result)
}

func main(){
	runApp(0)
}