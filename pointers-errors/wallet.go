package pointerserrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin float64

type Wallet struct {
	Balance Bitcoin
}

// GO works with object copies!
// so if we modify a balance we are not modifying the actual object
// To do that we need to pass the actual pointer as reference...

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.Balance += amount
}

// pointer references especified with *
func (wallet *Wallet) GetBalance() Bitcoin {
	return wallet.Balance
}

func (wallet *Wallet) Whitdraw(amount Bitcoin) error {
	if wallet.Balance < amount {
		return ErrInsufficientFunds
	}

	wallet.Balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}
