package usecase

import (
	"reflect"
	"testing"

	"github.com/jefersonc/banking-go/src/test/mock"
)

func TestSuccessCreateAccount(t *testing.T) {
	t.Run("CreateAccount success test", func(t *testing.T) {
		repository := mock.NewMockedAccountRepository()
		uc := NewCreateAccount(repository)

		payload := CreateAccountPayload{
			DocumentNumber: mock.SuccessAccountDocument,
		}

		response, err := uc.Execute(payload)

		if nil != err ||
			payload.DocumentNumber != response.DocumentNumber {
			t.Errorf("Error when create account usecase")
		}
	})
}

func TestCreateAccountWithInvalidDocument(t *testing.T) {
	t.Run("CreateAccount with invalid document test", func(t *testing.T) {
		repository := mock.NewMockedAccountRepository()
		uc := NewCreateAccount(repository)

		payload := CreateAccountPayload{
			DocumentNumber: "1234567890",
		}

		response, err := uc.Execute(payload)

		if nil == err ||
			response != nil ||
			reflect.TypeOf(err).String() != "*usecase.UserError" {
			t.Errorf("Error when create account usecase")
		}
	})
}

func TestCreateAccountWithRepositoryFail(t *testing.T) {
	t.Run("CreateAccount with repository fail test", func(t *testing.T) {
		repository := mock.NewMockedAccountRepository()
		uc := NewCreateAccount(repository)

		payload := CreateAccountPayload{
			DocumentNumber: "49762668006",
		}

		response, err := uc.Execute(payload)

		if nil == err ||
			response != nil ||
			reflect.TypeOf(err).String() != "*usecase.ApplicationError" {
			t.Errorf("Error when create account usecase")
		}
	})
}
