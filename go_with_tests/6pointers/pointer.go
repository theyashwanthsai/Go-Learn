package pointers

import "fmt"


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

func (w *Wallet) Withdraw(amount Bitcoin){
	w.Amount -= amount
}