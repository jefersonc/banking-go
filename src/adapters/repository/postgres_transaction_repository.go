package repository

import (
	"fmt"
	"time"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/infra/database"
	"github.com/jefersonc/banking-go/src/vo"
)

type PosgresTransactionRepository struct {
	Client *database.PostgresClient
}

var (
	fetchTransactionsByAccountQuery = "SELECT t.id, t.operation_id, t.amount, t.created_at, o.description, o.finality FROM transactions t LEFT JOIN operations o ON t.operation_id = o.id WHERE t.account_id = '%s'"
	pushTransactionQuery            = "INSERT INTO transactions (id, operation_id, account_id, amount, created_at) VALUES ('%s', '%s', '%s', %f, '%s')"
)

func (p PosgresTransactionRepository) FetchByAccount(account *domain.Account) ([]*domain.Transaction, error) {
	rows, err := p.Client.Connection.Query(fmt.Sprintf(fetchTransactionsByAccountQuery, account.GetID().Value()))

	if err != nil {
		return nil, err
	}

	transactions := []*domain.Transaction{}

	for rows.Next() {
		var (
			ID          string
			operationID string
			amount      float64
			createdAt   string
			description string
			finality    string
		)

		rows.Scan(&ID, &operationID, &amount, &createdAt, &description, &finality)

		operationIDVo, _ := vo.NewID(operationID)
		operation := domain.NewOperation(operationIDVo, description, finality)
		amountVo, _ := vo.NewAmount(amount)
		transactionID, _ := vo.NewID(ID)
		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		transaction := domain.NewTransaction(transactionID, account, operation, amountVo, date)

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (p PosgresTransactionRepository) Push(transaction *domain.Transaction) error {
	_, err := p.Client.Connection.Exec(fmt.Sprintf(
		pushTransactionQuery,
		transaction.GetID().Value(),
		transaction.GetOperation().GetID().Value(),
		transaction.GetAccount().GetID().Value(),
		transaction.GetAmount().Value(),
		transaction.GetDate().Format("2006-01-02 15:04:05")))

	if err != nil {
		return err
	}

	return nil
}

func CreatePostgresTransactionRepository(client *database.PostgresClient) PosgresTransactionRepository {
	return PosgresTransactionRepository{
		Client: client,
	}
}
