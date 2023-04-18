package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T){

	channel := make(chan string)
	defer close(channel)

	go func() {
		// time.Sleep(2 * time.Second)
		channel <- "Taylor alison swift"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}