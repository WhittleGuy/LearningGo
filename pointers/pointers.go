package main

import (
	"errors"
	"fmt"
)

// Defined in "fmt". Allows us to decide how %s is displayed when used with type
type Stringer interface {
	String() string
}

type Etherium float64

// Stringer def for Etherium
func (e Etherium) String() string {
	return fmt.Sprintf("%f ETH", e)
}

type Wallet struct {
	balance Etherium // Lowercase makes private
}

func (w *Wallet) Deposit(amount Etherium) { // * points to wallet mem, not copied var
	w.balance += amount
}

func (w *Wallet) Balance() Etherium { // * points to wallet mem, not copied var
	return w.balance
}

var ErrInsufficientFunds = errors.New("Insufficient funds")

func (w *Wallet) Withdraw(amount Etherium) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
