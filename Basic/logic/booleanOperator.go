package main

import "fmt"

func main(){
	var ujian = 80
	var absensi = 60

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi > 59

	var dosenKiller = lulusUjian && lulusAbsensi
	var dosenAsik = lulusAbsensi || lulusUjian

	fmt.Println("*",lulusUjian && lulusAbsensi)

	fmt.Println(dosenKiller)
	fmt.Println(dosenAsik)
}