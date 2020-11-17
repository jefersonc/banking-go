package repository

import (
	"fmt"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/infra/database"
	"github.com/jefersonc/banking-go/src/vo"
)

type PostgresOperationRepository struct {
	Client *database.PostgresClient
}

var (
	findOperationQuery = "SELECT description, finality FROM operations WHERE id = '%s'"
)

func (p PostgresOperationRepository) Find(id *vo.ID) (*domain.Operation, error) {
	var (
		description string
		finality    string
	)

	err := p.Client.Connection.QueryRow(fmt.Sprintf(findOperationQuery, id.Value())).Scan(&description, &finality)

	if err != nil {
		return nil, err
	}

	return domain.NewOperation(id, description, finality), nil
}

func CreatePostgresOperationRepository(client *database.PostgresClient) PostgresOperationRepository {
	return PostgresOperationRepository{
		Client: client,
	}
}
