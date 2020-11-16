package usecase

import (
	"testing"

	"github.com/jefersonc/banking-go/src/test/mock"
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
