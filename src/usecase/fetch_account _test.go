package usecase

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type mockedAccountRepository struct{}

var (
	successID = vo.GenerateID()
)

func (m mockedAccountRepository) Find(id *vo.ID) (*domain.Account, error) {
	if id.Value() == successID.Value() {
		document, _ := vo.NewDocument("CPF", "80987246038")
		account := domain.NewAccount(successID, document)

		return account, nil
	}

	return nil, errors.New("mocked")
}

func (m mockedAccountRepository) Push(account *domain.Account) error {
	if account.GetID().Value() == successID.Value() {
		return nil
	}

	return errors.New("mocked")
}

func TestSucessUseCase(t *testing.T) {
	t.Run("FetchAccount success test", func(t *testing.T) {
		repository := mockedAccountRepository{}
		uc := NewFetchAccount(repository)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080/accounts/%s", successID.Value()), nil)

		uc.Invoke(w, req)
	})
}
