package main

import "fmt"

func factorialLoop1(value int) int {
	result := 1
	for i := value;i > 0; i--{
		result *= i
	}
	return result
}

func factorialLoop2(value int) int {
	if value == 1 {
		return 1
	}else{
		return value * factorialLoop2(value - 1)
	}
}

func main(){
	fmt.Println(5*4*3*2*1)
	fmt.Println(factorialLoop2(5))


}