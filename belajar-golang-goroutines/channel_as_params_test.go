package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "taylor swift"
}


func TestChannelAsParams(t *testing.T){
	channel := make(chan string)
	defer close(channel)
	
	go GiveMeResponse(channel)

	fmt.Println("lalalala")

	data := <- channel

	fmt.Println(data)
}

func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)

	channel <- "Rosseane park"
}

func OnlyOut(channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOut(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	fmt.Println("gue dulu dong !!!")

	time.Sleep(3 * time.Second)

}


func TestBufferedChannel(t *testing.T){
	channel := make(chan string,3)
	defer close(channel)

	go func(){
		channel <- "taylor"
		channel <- "alison"
		channel <- "swift"
	}()

	go func(){
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(1 * time.Second)
}

func TestRangeChannel(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i < 10;i++{
			channel <- "perlulang ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("menerima data ke ",data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select{
		case data := <-channel1:
			fmt.Println("data dari channel 1",data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2",data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

func TestSelectDefaultChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select{
		case data := <-channel1:
			fmt.Println("data dari channel 1",data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2",data)
			counter++
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}

}


