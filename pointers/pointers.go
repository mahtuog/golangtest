package pointers

import (
	"errors"
	"fmt"
)

type BitCoin int

var ErrInsufficientFunds = errors.New("Cannot Withdraw, Insufficient Funds")

type Stringer interface {
	String() string
}

type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount BitCoin){
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin{
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) error{

	if amount > w.balance{
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}