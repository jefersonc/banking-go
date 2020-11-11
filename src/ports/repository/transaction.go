package repository

import (
	"github.com/jefersonc/banking-go/src/domain/entity"
)

// Transaction repository manage all interactions on table account
type Transaction interface {
	Push(transaction entity.Transaction) bool
}
