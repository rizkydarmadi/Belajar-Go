package BelajarGolangContext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestContextwithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))

	fmt.Println(contextG.Value("g"))
	fmt.Println(contextG.Value("f"))
	fmt.Println(contextG.Value("c"))

}

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			destination <- counter
			counter++
		}

	}()

	return destination
}
