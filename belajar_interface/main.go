package main

import "fmt"

type Calculator interface {
	Perkalian() int
	Penjumlahan() int
}

type Nilai struct {
	a, b int
}

func (n Nilai) Perkalian() int {
	return n.a * n.b
}

func (n Nilai) Penjumlahan() int {
	return n.a + n.b
}

func main() {

	var c Calculator
	c = Nilai{2, 3}

	fmt.Println(c.Perkalian())
	fmt.Println(c.Penjumlahan())
}
