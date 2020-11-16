package domain

import (
	"testing"
	"time"

	"github.com/jefersonc/banking-go/src/vo"
)

func TestSuccessNewTransaction(t *testing.T) {

	t.Run("Transaction instance test", func(t *testing.T) {
		accountID := vo.GenerateID()
		document, _ := vo.NewDocument("CPF", "42145390014")
		account := NewAccount(accountID, document)

		operationID := vo.GenerateID()
		description := "Depósito"
		operationType := OperationCredit
		operation := NewOperation(operationID, description, operationType)

		transactionID := vo.GenerateID()
		amount, _ := vo.NewAmount(10.0)
		date := time.Now()
		transaction := NewTransaction(transactionID, account, operation, amount, date)

		if transaction.GetID().Value() != transactionID.Value() ||
			transaction.GetAccount().GetID().Value() != accountID.Value() ||
			transaction.GetOperation().GetID().Value() != operationID.Value() ||
			transaction.GetAmount().Value() != amount.Value() ||
			transaction.GetDate().String() != date.String() {
			t.Errorf("Error when checking transaction instance")
		}
	})
}
