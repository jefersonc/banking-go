package domain

import (
	"github.com/jefersonc/banking-go/src/vo"
)

// AccountRepository repository manage all interactions with persistence
type AccountRepository interface {
	Find(id *vo.ID) (*Account, error)
	Push(account *Account) error
}

// Account is entity/aggregator
type Account struct {
	id       *vo.ID
	document *vo.Document
}

// GetID is a getter for document_number attribute
func (a Account) GetID() *vo.ID {
	return a.id
}

// GetDocument is a getter for document_number attribute
func (a Account) GetDocument() *vo.Document {
	return a.document
}

// AccountFactory is a constructor
func NewAccount(id *vo.ID, document *vo.Document) *Account {
	return &Account{id, document}
}
