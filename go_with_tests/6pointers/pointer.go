package pointers

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Stringer interface{
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}
type Wallet struct{
	Amount Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin){
	w.Amount += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.Amount
}


var InsufficientFundsErr = errors.New("cannot withdraw, insufficient funds")
func (w *Wallet) Withdraw(amount Bitcoin) error{
	if w.Amount < amount{
		return InsufficientFundsErr
	}else{
		w.Amount -= amount
	}
	return nil

}