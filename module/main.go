package main

import (
	"fmt"
	"go-modules/calculator"
	"go-modules/persegi"
)

func main() {
	fmt.Println("test")

	var c calculator.HitungCalculator
	c = calculator.Nilai{2, 2, 2}

	fmt.Println(c.Perkalian())

	var p persegi.Persegi
	p = persegi.Persegi{5}

	fmt.Println(p.Luas())

}
