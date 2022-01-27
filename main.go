package main

import (
	"banking/banking"
	"fmt"
	"log"
)

//banking account export 방법.

func main() {
	account := banking.NewAccount("js")	
	account.Deposit(10)	
	err := account.Withdraw(20)

	if err != nil {
		log.Println(err)
		//fatal은 프로그램종료 역할
	} 
	fmt.Println(account.Balance(), account.Owner(), account)
}