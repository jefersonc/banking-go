package repository

import (
	"time"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type PosgresTransactionRepository struct{}

func (p PosgresTransactionRepository) FetchByAccount(account *domain.Account) ([]*domain.Transaction, error) {

	id := vo.GenerateID()

	operation := domain.NewOperation(id, "teste", "CREDIT")
	amount, _ := vo.NewAmount(10.00)
	transaction := domain.NewTransaction(id, account, operation, amount, time.Now())

	transactions := []*domain.Transaction{
		transaction,
	}

	return transactions, nil
}

func (p PosgresTransactionRepository) Push(transaction *domain.Transaction) error {
	return nil
}

func CreatePostgresTransactionRepository() PosgresTransactionRepository {
	return PosgresTransactionRepository{}
}
