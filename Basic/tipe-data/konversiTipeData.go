package main

import "fmt"

func main(){
	// konversi tipe data angka
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) //nilai akan kembali ke titik terbawah karena int8 nilai maksimalnya hanyalah 127

	//konversi byte to string
	var name = "ariana grande"
	var e = name[0]
	var convertString = string(e)

	fmt.Println(name) 
	fmt.Println(e)
	fmt.Println(convertString)
}
