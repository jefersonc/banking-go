package repository

import (
	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type PosgresAccountRepository struct{}

func (p PosgresAccountRepository) Find(id *vo.ID) (*domain.Account, error) {
	document, _ := vo.NewDocument("CPF", "08732860900")
	return domain.NewAccount(id, document), nil
}

func (p PosgresAccountRepository) Push(account *domain.Account) error {
	return nil
}

func CreatePostgresAccountRepository() PosgresAccountRepository {
	return PosgresAccountRepository{}
}
