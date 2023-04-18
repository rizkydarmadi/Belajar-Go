package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T){
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	
	group := sync.WaitGroup{}

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0;i<10;i++{
		go func ()  {
			defer group.Done()
			group.Add(1)

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	group.Wait()
}