package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(sum float64) {
	b.balance += sum
}

func (b *BankAccount) Withdraw(sum float64) error {
	if b.balance < sum {
		return errors.New("Недостаточно средств")
	}
	b.balance -= sum
	return nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func main() {
	account := &BankAccount{
		owner:   "Иван",
		balance: 1000,
	}

	account.Deposit(500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())

	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())
	}

	err = account.Withdraw(1500)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())
	}
}
