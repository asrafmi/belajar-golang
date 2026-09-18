package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 + time.Second)
	fmt.Println("Counter: ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance: ", account.Balance) // output: 10000 (100x100)
}

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amoun int) {
	user.Balance = user.Balance + amoun
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func SuccessTransfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	// selalu lock yang "lebih kecil" duluan, siapapun yang jadi user1/user2
	first, second := user1, user2
	if second.Name < first.Name {
		first, second = second, first
	}

	first.Lock()
	fmt.Println("Lock", first.Name)
	defer first.Unlock()

	time.Sleep(1 * time.Second)

	second.Lock()
	fmt.Println("Lock", second.Name)
	defer second.Unlock()

	user1.Change(-amount)
	user2.Change(amount)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Budi",
		Balance: 2000000,
	}

	// go Transfer(&user1, &user2, 1000)
	// go Transfer(&user2, &user1, 2000)
	go SuccessTransfer(&user1, &user2, 1000)
	go SuccessTransfer(&user2, &user1, 2000)

	time.Sleep(3 * time.Second)

	fmt.Println("User 1: ", user1.Balance)
	fmt.Println("User 2: ", user2.Balance)
}
