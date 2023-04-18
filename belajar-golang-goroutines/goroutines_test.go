package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(){
	fmt.Println("hello world")
}

func TestCreateGoroutines(t *testing.T){
	go RunHelloWorld()
	fmt.Println("ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("number",number)
}


func TestManyNumber(t *testing.T){
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}



