package main

func usePooledAccountFixed() int {
	account := pool.Get().(*Account)
	defer pool.Put(account)

	account.total = 0

	account.total += 5

	return account.total
}