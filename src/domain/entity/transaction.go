package entity

import (
	"time"
)

// Transaction is entity/aggregator
type Transaction struct {
	id            int
	account       Account
	operationType OperationType
	amount        float32
	date          time.Time
}

// GetID is a getter for id attribute
func (t Transaction) GetID() int {
	return t.id
}

// GetAccount is a getter for account attribute
func (t Transaction) GetAccount() Account {
	return t.account
}

// GetOperationType is a getter for operationType attribute
func (t Transaction) GetOperationType() OperationType {
	return t.operationType
}

// GetAmount is a getter for amount attribute
func (t Transaction) GetAmount() float32 {
	return t.account
}

// GetDate is a getter for date attribute
func (t Transaction) GetDate() time.Time {
	return t.date
}

// CreateTransaction is a constructor
func CreateTransaction(
	id int,
	account Account,
	operationType OperationType,
	amount float32,
	date time.Time) Transaction {
	return Transaction{id, account, operationType, amount, date}
}
