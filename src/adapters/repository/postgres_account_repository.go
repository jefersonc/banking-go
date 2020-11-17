package repository

import (
	"fmt"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/infra/database"
	"github.com/jefersonc/banking-go/src/vo"
)

type PosgresAccountRepository struct {
	Client *database.PostgresClient
}

var (
	findAccountQuery = "SELECT document_type, document_number FROM accounts WHERE id = '%s'"
	pushAccountQuery = "INSERT INTO accounts (id, document_type, document_number) VALUES ('%s', '%s', '%s')"
)

func (p PosgresAccountRepository) Find(id *vo.ID) (*domain.Account, error) {
	var (
		document_type   string
		document_number string
	)

	err := p.Client.Connection.QueryRow(fmt.Sprintf(findAccountQuery, id.Value())).Scan(&document_type, &document_number)

	if err != nil {
		return nil, err
	}

	document, _ := vo.NewDocument(document_type, document_number)

	return domain.NewAccount(id, document), nil
}

func (p PosgresAccountRepository) Push(account *domain.Account) error {
	_, err := p.Client.Connection.Exec(fmt.Sprintf(pushAccountQuery, account.GetID().Value(), account.GetDocument().Type(), account.GetDocument().Number()))
	if err != nil {
		return err
	}

	return nil
}

func CreatePostgresAccountRepository(client *database.PostgresClient) *PosgresAccountRepository {
	return &PosgresAccountRepository{
		Client: client,
	}
}
