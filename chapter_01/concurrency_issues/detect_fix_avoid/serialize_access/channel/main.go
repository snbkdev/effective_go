package main

func NewAccount() *Account {
	account := &Account{
		actionsCh: make(chan *action),
	}

	go account.processActions()

	return account
}

type Account struct {
	actionsCh chan *action
}

func (a *Account) Transfer(amount int, destination *Account) {
	withdrawAction := newAction(-1 * amount)
	a.actionsCh <- withdrawAction

	destination.Deposit(amount)
}

func (a *Account) Deposit(amount int) {
	action := newAction(amount)

	a.actionsCh <- action
}

func (a *Account) Balance() int {
	action := newAction(0)

	a.actionsCh <- action

	return <- action.resultCh
}

func (a *Account) processActions() {
	var balance int

	for thisAction := range a.actionsCh {
		balance += thisAction.amount
		thisAction.resultCh <- balance
	}
}

func newAction(amount int) *action {
	return &action{
		amount: amount,
		resultCh: make(chan int, 1),
	}
}

type action struct {
	amount   int
	resultCh chan int
}