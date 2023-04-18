package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T){

	group := sync.WaitGroup{}
	for i:=0; i < 100;i++{
		group.Add(1)

		go func(){
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}


	totalCpu := runtime.NumCPU()
	fmt.Println("num cpu",totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread",totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine",totalGoroutine)

	group.Wait()
}