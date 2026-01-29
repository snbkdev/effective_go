package main

import "sync"

type Account struct {
	balance int
	mutex *sync.Mutex
}

func (a *Account) Transfer(amount int, destination *Account) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.balance -= amount

	destination.Deposit(amount)
}

func (a *Account) Deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.balance += amount
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	return a.balance
}