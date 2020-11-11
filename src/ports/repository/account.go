package repository

import (
	"github.com/jefersonc/banking-go/src/domain/entity"
)

// Account repository manage all interactions on table account
type Account interface {
	Find(id int) entity.Account
	Push(account entity.Account) bool
}
