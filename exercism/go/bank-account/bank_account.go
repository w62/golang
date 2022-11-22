package account

import (
	"sync"
)

// Define the Account type here.

type Account struct {
	mu     sync.Mutex
	name   string
	amount int64
	state  bool // true means open false means closed
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	var acc Account
	acc.name = ""
	acc.amount = amount
	acc.state = true
	return &acc
	panic("Please implement the Open function")
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if (*a).state == false {
		return 0, false
	}

	return (*a).amount, true
	panic("Please implement the Balance function")
}

func (a *Account) Deposit(amount int64) (int64, bool) {

	a.mu.Lock()
	defer a.mu.Unlock()
	if a.state == false {
		return 0, false
	}

	if amount <= 0 && amount > a.amount {
		return 0, false
	}

	if a.amount+amount < 0 {
		return 0, false
	}

	a.amount += amount
	return a.amount, true

	panic("Please implement the Deposit function")
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.state == false {
		a.amount = 0
		return 0, false
	}

	a.state = false
	return a.amount, true
	panic("Please implement the Close function")
}
