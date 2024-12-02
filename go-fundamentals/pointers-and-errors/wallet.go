package pointers_and_errors

import (
	"errors"
	"fmt"
)

// Bitcoin type - note could be fractioned
type Bitcoin float64

// Implement Stringer for BTC
func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

// Wallet type
type Wallet struct {
	balance Bitcoin
}

// Deposit BTC in Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

// Retrieve Balance from Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Error associated with insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw BTC in Wallet, can throw error if insufficient funds
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
