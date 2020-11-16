package domain

import (
	"github.com/jefersonc/banking-go/src/vo"
)

var (
	OperationCredit = "CREDIT"
	OperationDebit  = "DEBIT"
)

// OperationRepository repository manage all interactions with persistence
type OperationRepository interface {
	Find(id *vo.ID) (*Operation, error)
}

// Operation is entity/aggregator
type Operation struct {
	id          *vo.ID
	description string
	finality    string
}

// GetID is a getter for id attribute
func (o *Operation) GetID() *vo.ID {
	return o.id
}

// GetDescription is a getter for description attribute
func (o *Operation) GetDescription() string {
	return o.description
}

func (o *Operation) GetFinality() string {
	return o.finality
}

// NewOperation is a constructor
func NewOperation(id *vo.ID, description string, finality string) *Operation {
	return &Operation{id, description, finality}
}
