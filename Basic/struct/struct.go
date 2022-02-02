package main

import "fmt"

type Coustemer struct {
	Name, Addres string
	Age          int
}

func (coustemer Coustemer) sayHi(name string) {
	fmt.Println("hello", name, "my name is", coustemer.Name)
}

func main() {
	var biodata Coustemer
	biodata.Name = "rizky"
	biodata.Addres = "indonesia"
	biodata.Age = 23

	fmt.Println(biodata)

	bowo := Coustemer{
		Name:   "bowo",
		Addres: "Indonesia",
		Age:    30,
	}

	fmt.Println(bowo)

	/**
	untuk struct seperti ini perlu urutanya diingiat agar tidak berarturan wajib
	sesuai dengan deklarasi pertamanya diatas
	*/

	julian := Coustemer{
		"julian",
		"indonesia",
		28,
	}
	fmt.Println(julian)

	bowo.sayHi("marisa")

}
