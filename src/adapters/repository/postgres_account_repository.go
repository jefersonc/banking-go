package repository

import (
	"github.com/jefersonc/banking-go/src/domain/entity"
)

type PosgresAccountRepository struct{}

func (p PosgresAccountRepository) Find(id int) entity.Account {
	return entity.CreateAccount(1, "84878787887")
}

func (p PosgresAccountRepository) Push(account entity.Account) bool {
	return true
}

func CreatePostgresAccountRepository() PosgresAccountRepository {
	return PosgresAccountRepository{}
}
