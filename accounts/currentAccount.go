package accounts

import (
	"go-bank/client"
)

type CurrentAccount struct {
	Holder                      client.Holder
	NumberAccount, NumberAgency int
	balance                     float64
}

func (account *CurrentAccount) Withdraw(value float64) (string, bool) {
	if value > 0 && account.balance > value {
		account.balance -= value
		return "Withdrawal successful", true
	}
	return "Insufficient funds", false
}

func (account *CurrentAccount) Deposit(value float64) (string, bool) {
	if value > 0 {
		account.balance += value
		return "Deposit successful", true
	}
	return "Invalid deposit value", false
}

func (account *CurrentAccount) Transfer(value float64, destinationAccount *CurrentAccount) (string, bool) {
	message, status := account.Withdraw(value)
	if status {
		message, status := destinationAccount.Deposit(value)
		if status {
			return "Transfer successful", true
		} else {
			account.Deposit(value)
			return message, status
		}
	}
	return message, status
}

func (account *CurrentAccount) GetBalance() float64 {
	return account.balance
}
