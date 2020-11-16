package usecase

import (
	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type (
	CreateAccount struct {
		repository domain.AccountRepository
	}

	CreateAccountPayload struct {
		DocumentNumber string `json:"document_number"`
	}

	CreateAccountResponse struct {
		ID             string `json:"id"`
		DocumentType   string `json:"document_type"`
		DocumentNumber string `json:"document_number"`
	}
)

func (uc *CreateAccount) Execute(payload CreateAccountPayload) (*CreateAccountResponse, error) {
	document, err := vo.NewDocument("CPF", payload.DocumentNumber)

	if err != nil {
		return nil, NewUserError(err)
	}

	account := domain.NewAccount(vo.GenerateID(), document)

	err = uc.repository.Push(account)

	if err != nil {
		return nil, NewApplicationError(err)
	}

	return uc.output(account), nil
}

func (uc *CreateAccount) output(account *domain.Account) *CreateAccountResponse {
	return &CreateAccountResponse{
		ID:             account.GetID().Value(),
		DocumentType:   account.GetDocument().Type(),
		DocumentNumber: account.GetDocument().Number(),
	}
}

func NewCreateAccount(repository domain.AccountRepository) CreateAccount {
	return CreateAccount{repository}
}
