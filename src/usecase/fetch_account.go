package usecase

import (
	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type (
	FetchAccount struct {
		repository domain.AccountRepository
	}

	FetchAccountPayload struct {
		ID string
	}

	FetchAccountResponse struct {
		ID             string `json:"id"`
		DocumentType   string `json:"document_type"`
		DocumentNumber string `json:"document_number"`
	}
)

func (uc *FetchAccount) Execute(payload FetchAccountPayload) (*FetchAccountResponse, error) {
	id, err := vo.NewID(payload.ID)

	if err != nil {
		return nil, NewUserError(err)
	}

	account, err := uc.repository.Find(id)

	if err != nil {
		return nil, NewApplicationError(err)
	}

	return uc.output(account), nil
}

func (uc *FetchAccount) output(account *domain.Account) *FetchAccountResponse {
	return &FetchAccountResponse{
		ID:             account.GetID().Value(),
		DocumentType:   account.GetDocument().Type(),
		DocumentNumber: account.GetDocument().Number(),
	}
}

func NewFetchAccount(repository domain.AccountRepository) FetchAccount {
	return FetchAccount{repository}
}
