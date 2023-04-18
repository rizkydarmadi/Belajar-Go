package calculator

type CalculatorInt interface {
	Penjumlahan() int
	Pembagian() int
	Perkalian() int
}

type Values struct {
	A, B int
}

func (n Values) Penjumlahan() int {
	return n.A + n.B
}

func (n Values) Perkalian() int {
	return n.A * n.B
}

func (n Values) Pembagian() int {
	return n.A / n.B
}
