package main

import "sync"

var pool = &sync.Pool{
	New: func() interface{} {
		return &Account{}
	},
}

func usePooledAccount() int {
	account := pool.Get().(*Account)
	defer pool.Put(account)

	account.total += 5

	return account.total
}

type Account struct {
	total int
}