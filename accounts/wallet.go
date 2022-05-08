package accounts

import "errors"

var nsfError = errors.New("NSF")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdrawl(amount Bitcoin) error {
	if amount > w.balance {
		return nsfError
	}
	w.balance -= amount
	return nil
}
