package mock

import (
	"errors"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

var (
	SuccessAccountID       = vo.GenerateID()
	SuccessAccountDocument = "80987246038"

	CreditOperationID = vo.GenerateID()
	DebitOperaionID   = vo.GenerateID()
	FailOperationID   = vo.GenerateID()
)

type (
	MockedAccountRepository     struct{}
	MockedOperationRepository   struct{}
	MockedTransactionRepository struct{}
)

func (m *MockedAccountRepository) Find(id *vo.ID) (*domain.Account, error) {
	if id.Value() == SuccessAccountID.Value() {
		document, _ := vo.NewDocument("CPF", "80987246038")
		account := domain.NewAccount(SuccessAccountID, document)

		return account, nil
	}

	return nil, errors.New("mocked")
}

func (m *MockedAccountRepository) Push(account *domain.Account) error {
	if account.GetDocument().Number() == SuccessAccountDocument {
		return nil
	}

	return errors.New("mocked")
}

func NewMockedAccountRepository() *MockedAccountRepository {
	return &MockedAccountRepository{}
}

func (m *MockedOperationRepository) Find(id *vo.ID) (*domain.Operation, error) {
	if id.Value() == CreditOperationID.Value() {
		return domain.NewOperation(id, "teste", domain.OperationCredit), nil
	}
	if id.Value() == DebitOperaionID.Value() {
		return domain.NewOperation(id, "teste", domain.OperationDebit), nil
	}

	return nil, errors.New("mocked")
}

func NewMockedOperationRepository() *MockedOperationRepository {
	return &MockedOperationRepository{}
}
