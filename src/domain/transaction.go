package domain

import (
	"time"

	"github.com/jefersonc/banking-go/src/vo"
)

// TransactionRepository repository manage all interactions with persistence layer
type TransactionRepository interface {
	Push(transaction *Transaction) error
	FetchByAccount(account *Account) (*[]Transaction, error)
}

// Transaction is entity/aggregator
type Transaction struct {
	id        *vo.ID
	account   *Account
	operation *Operation
	amount    *vo.Amount
	date      time.Time
}

// GetID is a getter for id attribute
func (t Transaction) GetID() *vo.ID {
	return t.id
}

// GetAccount is a getter for account attribute
func (t Transaction) GetAccount() *Account {
	return t.account
}

// Getoperation is a getter for operation attribute
func (t Transaction) Getoperation() *Operation {
	return t.operation
}

// GetAmount is a getter for amount attribute
func (t Transaction) GetAmount() *vo.Amount {
	return t.amount
}

// GetDate is a getter for date attribute
func (t Transaction) GetDate() time.Time {
	return t.date
}

// CreateTransaction is a constructor
func NewTransaction(
	id *vo.ID,
	account *Account,
	operation *Operation,
	amount *vo.Amount,
	date time.Time) *Transaction {
	return &Transaction{id, account, operation, amount, date}
}
