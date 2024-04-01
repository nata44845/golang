package main

import (
	"errors"
	"fmt"
	"time"
)

type Operator interface {
	Balance() int64
	Withdraw(amount int64) error
	Deposit(amount int64) error
	Transactions() []Tx
}

type ActionKind string

const (
	ActionKindIncr ActionKind = "+"
	ActionKindDecr ActionKind = "-"
)

type Tx struct {
	value     int64      // значение на которое изменилось
	action    ActionKind // действие, прибавляем или отнимаем
	createdAt time.Time
}

// Нужно вывести данные транзакции в формате сумма: +-value, time: время создания транзакции
func (t Tx) Print() string {
	return fmt.Sprintf("сумма: %s%d, time: %s", t.action, t.value, t.createdAt.Format(time.DateTime))
}

var _ Operator = (*txOperator)(nil)

type txOperator struct {
	balance      int64
	transactions []Tx
}

func (t *txOperator) Balance() int64 {
	// TODO implement me
	return t.balance
	// panic("implement me")
}

func (t *txOperator) Withdraw(amount int64) error {
	// TODO implement me
	if amount > t.balance {
		return errors.New("Недостаточно средств")
	}
	t.balance -= amount
	t.transactions = append(t.transactions, Tx{amount, ActionKindDecr, time.Now()})
	return nil
}

func (t *txOperator) Deposit(amount int64) error {
	t.balance += amount
	t.transactions = append(t.transactions, Tx{amount, ActionKindIncr, time.Now()})
	return nil
}

func (t *txOperator) Transactions() []Tx {
	// TODO implement me
	return t.transactions
}

func main() {
	var op Operator = &txOperator{}
	_ = op
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
	op.Deposit(100)
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())

	for _, t := range op.Transactions() {
		fmt.Println(t.Print())
	}

}
