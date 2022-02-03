package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "anthony",
		"addres": "jl Fatmawati",
	}
	person["title"] = "employee"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	var book map[string]string = make(map[string]string)
	book["title"] = "filososfi teras"
	book["author"] = "jk rowlink"
	book["ups"] = "salah"

	delete(book, "ups")

	fmt.Println(book)

}
