package main

import (
	"fmt"
	"latihan1/calculator"
)

func main() {
	var c calculator.CalculatorInt
	c = calculator.Values{2, 4}

	fmt.Println(c.Penjumlahan())

}
