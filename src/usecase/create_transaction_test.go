package usecase

import (
	"reflect"
	"testing"

	"github.com/jefersonc/banking-go/src/test/mock"
	"github.com/jefersonc/banking-go/src/vo"
)

func TestSuccessCreateTransaction(t *testing.T) {
	t.Run("CreateTransaction success test", func(t *testing.T) {
		uc := NewCreateTransaction(mock.NewMockedAccountRepository(), mock.NewMockedOperationRepository(), mock.NewMockedTransactionRepository())

		payload := CreateTransactionPayload{
			AccountID:       mock.SuccessAccountID.Value(),
			OperationTypeID: mock.CreditOperationID.Value(),
			Amount:          10.00,
		}

		_, err := uc.Execute(payload)

		if nil != err {
			t.Errorf("Error when create transaction usecase")
		}
	})
}

func TestCreateTransactionWithInvalidAccount(t *testing.T) {
	t.Run("CreateTransaction with invalid account", func(t *testing.T) {
		uc := NewCreateTransaction(mock.NewMockedAccountRepository(), mock.NewMockedOperationRepository(), mock.NewMockedTransactionRepository())

		payload := CreateTransactionPayload{
			AccountID:       vo.GenerateID().Value(),
			OperationTypeID: mock.CreditOperationID.Value(),
			Amount:          10.00,
		}

		_, err := uc.Execute(payload)

		if nil == err ||
			reflect.TypeOf(err).String() != "*usecase.ApplicationError" {
			t.Errorf("Error when create transaction usecase")
		}
	})
}

func TestCreateTransactionWithInvalidOperation(t *testing.T) {
	t.Run("CreateTransaction with invalid operation", func(t *testing.T) {
		uc := NewCreateTransaction(mock.NewMockedAccountRepository(), mock.NewMockedOperationRepository(), mock.NewMockedTransactionRepository())

		payload := CreateTransactionPayload{
			AccountID:       mock.SuccessAccountID.Value(),
			OperationTypeID: vo.GenerateID().Value(),
			Amount:          10.00,
		}

		_, err := uc.Execute(payload)

		if nil == err ||
			reflect.TypeOf(err).String() != "*usecase.ApplicationError" {
			t.Errorf("Error when create transaction usecase")
		}
	})
}

func TestCreateTransactionWithInvalidAmount(t *testing.T) {
	t.Run("CreateTransaction with invalid amount", func(t *testing.T) {
		uc := NewCreateTransaction(mock.NewMockedAccountRepository(), mock.NewMockedOperationRepository(), mock.NewMockedTransactionRepository())

		payload := CreateTransactionPayload{
			AccountID:       mock.SuccessAccountID.Value(),
			OperationTypeID: mock.CreditOperationID.Value(),
			Amount:          -10.00,
		}

		_, err := uc.Execute(payload)

		if nil == err ||
			reflect.TypeOf(err).String() != "*usecase.UserError" {
			t.Errorf("Error when create transaction usecase")
		}
	})
}
