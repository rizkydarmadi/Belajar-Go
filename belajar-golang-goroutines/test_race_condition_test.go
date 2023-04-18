package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T){
	counter := 1

	for i := 0;i<1000;i++{
		go func(){
			for j := 0;j<100;j++{
				counter += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(counter)
}

func TestRaceConditionMutexCondition(t *testing.T){
	counter := 0

	var mutex sync.Mutex

	for i := 0;i<1000;i++{
		go func(){
			
			for j := 0;j<100;j++{
				mutex.Lock()
				counter += 1
				mutex.Unlock()
			}
			
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(counter)
}



type BankAccount struct{
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) addBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) getBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T){
	account := BankAccount{}

	for i:=0;i<1000;i++{
		go func(){
			for j :=0; j < 100; j++ {
				account.addBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}

	time.Sleep(10 * time.Second)
}