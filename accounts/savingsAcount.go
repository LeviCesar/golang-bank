package accounts

import (
	"go-bank/client"
)

type SavingsAccount struct {
	Holder                                 client.Holder
	NumberAccount, NumberAgency, Operation int
	balance                                float64
}

func (account *SavingsAccount) Withdraw(value float64) (string, bool) {
	if value > 0 && account.balance > value {
		account.balance -= value
		return "Withdrawal successful", true
	}
	return "Insufficient funds", false
}

func (account *SavingsAccount) Deposit(value float64) (string, bool) {
	if value > 0 {
		account.balance += value
		return "Deposit successful", true
	}
	return "Invalid deposit value", false
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}
