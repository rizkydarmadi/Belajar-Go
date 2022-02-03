package main

import "fmt"

// interface kosong interface{} dapat mereturn tipe data apapun

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	// untuk menginisialisasikan interface kosong wajib memmanggilnya kembali saat inisialisasi
	var data interface{} = Ups(3)
	fmt.Println(data)
}
