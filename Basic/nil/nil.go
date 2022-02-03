package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

/**
Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
*/
func main() {
	var person map[string]string = newMap("")

	if person == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person)
	}
}
