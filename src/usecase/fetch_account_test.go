package usecase

import (
	"reflect"
	"testing"

	"github.com/jefersonc/banking-go/src/test/mock"
	"github.com/jefersonc/banking-go/src/vo"
)

func TestSuccessFetchAccount(t *testing.T) {
	t.Run("FetchAccount success test", func(t *testing.T) {
		repository := mock.NewMockedAccountRepository()
		uc := NewFetchAccount(repository)

		payload := FetchAccountPayload{
			ID: mock.SuccessAccountID.Value(),
		}

		response, err := uc.Execute(payload)

		if nil != err ||
			payload.ID != response.ID {
			t.Errorf("Error when fetch account usecase")
		}
	})
}

func TestFetchAccountWithInvalidId(t *testing.T) {
	t.Run("FetchAccount with invalid document test", func(t *testing.T) {
		repository := mock.NewMockedAccountRepository()
		uc := NewFetchAccount(repository)

		payload := FetchAccountPayload{
			ID: "asdfg",
		}

		response, err := uc.Execute(payload)

		if nil == err ||
			response != nil ||
			reflect.TypeOf(err).String() != "*usecase.UserError" {
			t.Errorf("Error when fetch account usecase")
		}
	})
}

func TestFetchAccountWithRepositoryFail(t *testing.T) {
	t.Run("FetchAccount with repository fail test", func(t *testing.T) {
		repository := mock.NewMockedAccountRepository()
		uc := NewFetchAccount(repository)

		payload := FetchAccountPayload{
			ID: vo.GenerateID().Value(),
		}

		response, err := uc.Execute(payload)

		if nil == err ||
			response != nil ||
			reflect.TypeOf(err).String() != "*usecase.ApplicationError" {
			t.Errorf("Error when fetch account usecase")
		}
	})
}
