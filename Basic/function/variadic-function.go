package main

import "fmt"

func sumAll(numbers ...int)int{
	total := 0
	for _, value := range numbers{
		total += value
	}
	return total
}

func main(){
	total := sumAll(4,4,4)
	fmt.Println(total)

	slice := []int{1,6,3}
	total = sumAll(slice...)
	fmt.Println(total)
}