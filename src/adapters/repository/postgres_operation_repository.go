package repository

import (
	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type PostgresOperationRepository struct{}

func (p PostgresOperationRepository) Find(id *vo.ID) (*domain.Operation, error) {
	return domain.NewOperation(id, "teste", "CREDIT"), nil
}

func CreatePostgresOperationRepository() PostgresOperationRepository {
	return PostgresOperationRepository{}
}
