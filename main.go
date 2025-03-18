package main

import (
	"fmt"
	"go-bank/accounts"
	"go-bank/client"
)

func PayBoleto(conta CheckAccount, value float64) {
	conta.Withdraw(value)
}

type CheckAccount interface {
	Withdraw(value float64) (string, bool)
}

func main() {
	holder := client.Holder{Name: "João", Document: "123.123.123-12", Profession: "Developer"}
	testAccount := accounts.CurrentAccount{Holder: holder, NumberAccount: 123, NumberAgency: 123}

	testAccount.Deposit(100)
	PayBoleto(&testAccount, 60)

	fmt.Println(testAccount.GetBalance())
}
