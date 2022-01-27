package banking

import (
	"errors"
	"fmt"
)

// Account struct
// type Account struct {
// 	Owner   string
// 	Balance int
// }

//대문자여야 다른곳에서 사용 가능. public.
// owner   string
// balance int
//소문자는 private라서 다른곳에서 못씀..

type Account struct {	
	owner   string
	balance int	
}

var errNoMoney = errors.New("Can't Withdraw")

func NewAccount(owner string) *Account{ //Account를 살펴봄..
	account := Account{owner: owner, balance: 0}
	return &account
}


//receiver
//Depodit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your Account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error{
	if a.balance < amount{
		return errNoMoney
	}	 
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner (newOwner string){
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}