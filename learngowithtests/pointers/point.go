package point

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

type Bitcoin int 


type Wallet struct {

	balance Bitcoin

}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}


func (w *Wallet) Deposit(amount Bitcoin){
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount

}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil

}

func (w *Wallet) Balance() Bitcoin {

	return w.balance

}


// ====== iter 1 simple balance

// type Wallet struct {

// 	balance int

// }

// func (w *Wallet) Deposit(amount int){
// 	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
// 	w.balance += amount

// }

// func (w *Wallet) Balance() int {

// 	return w.balance

// }