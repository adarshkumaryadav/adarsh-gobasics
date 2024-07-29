package pointerscode

import "fmt"

type Wallet struct {
	Amount int
}

func (w *Wallet) Deposit(amount int) int {
	w.Amount += amount
	return w.Amount
}

func (w *Wallet) Balace() int {
	return w.Amount
}

func (w *Wallet) Deduct(amount int) (remaining int, err error) {
	if amount > w.Amount {
		return 0, fmt.Errorf("you don't have enough balance to withdraw")
	}
	return w.Amount - amount, nil
}
