package domain

import (
	"testing"
	"time"

	"github.com/jefersonc/banking-go/src/vo"
)

func createMovimentationData() (*Account, []*Transaction) {
	accountID := vo.GenerateID()
	document, _ := vo.NewDocument("CPF", "42145390014")
	account := NewAccount(accountID, document)

	operationCreditID := vo.GenerateID()
	operationCreditDescription := "Depósito"
	operationCreditType := OperationCredit
	operationCredit := NewOperation(operationCreditID, operationCreditDescription, operationCreditType)

	operationDebitID := vo.GenerateID()
	operationDebitDescription := "Saque"
	operationDebitType := OperationDebit
	operationDebit := NewOperation(operationDebitID, operationDebitDescription, operationDebitType)

	transactionID := vo.GenerateID()
	amount, _ := vo.NewAmount(10.0)
	date := time.Now()

	transactions := []*Transaction{
		NewTransaction(transactionID, account, operationCredit, amount, date),
		NewTransaction(transactionID, account, operationCredit, amount, date),
		NewTransaction(transactionID, account, operationDebit, amount, date),
	}

	return account, transactions
}

func TestMovimentationInstance(t *testing.T) {

	t.Run("Movimentation instance test", func(t *testing.T) {
		account, transactions := createMovimentationData()

		movimentation := MovimentationFactory(account, transactions)

		if movimentation.GetAccount().GetID().Value() != account.GetID().Value() ||
			len(movimentation.GetTransactions()) != len(transactions) {
			t.Errorf("Error when checking movimentation instance")
		}
	})
}

func TestMovimentationBalance(t *testing.T) {

	t.Run("Movimentation balance test", func(t *testing.T) {
		account, transactions := createMovimentationData()

		movimentation := MovimentationFactory(account, transactions)
		var expected float64 = 10.0

		if expected != movimentation.GetBalance() {
			t.Errorf("Error when checking balance")
		}
	})
}

func TestMovimentationCheckMovimentationCredit(t *testing.T) {

	t.Run("Movimentation Check Movimentation test", func(t *testing.T) {
		account, transactions := createMovimentationData()

		movimentation := MovimentationFactory(account, transactions)

		operationCreditID := vo.GenerateID()
		operationCreditDescription := "Depósito"
		operationCreditType := OperationCredit
		operationCredit := NewOperation(operationCreditID, operationCreditDescription, operationCreditType)

		transactionID := vo.GenerateID()
		amount, _ := vo.NewAmount(10.0)
		date := time.Now()
		transactionIntention := NewTransaction(transactionID, account, operationCredit, amount, date)

		check := movimentation.CheckMovimentation(transactionIntention)

		if check != nil {
			t.Errorf("Error when checking movimentation")
		}
	})
}

func TestMovimentationCheckMovimentationDebitWithNoFounds(t *testing.T) {

	t.Run("Movimentation Check Movimentation test", func(t *testing.T) {
		account, transactions := createMovimentationData()

		movimentation := MovimentationFactory(account, transactions)

		operationDebitID := vo.GenerateID()
		operationDebitDescription := "Saque"
		operationDebitType := OperationDebit
		operationDebit := NewOperation(operationDebitID, operationDebitDescription, operationDebitType)

		transactionID := vo.GenerateID()
		amount, _ := vo.NewAmount(100.0)
		date := time.Now()
		transactionIntention := NewTransaction(transactionID, account, operationDebit, amount, date)

		check := movimentation.CheckMovimentation(transactionIntention)

		if check != ErrInsufficientFunds {
			t.Errorf("Error when checking movimentation")
		}
	})
}

func TestMovimentationCheckMovimentationDebitWithFounds(t *testing.T) {

	t.Run("Movimentation Check Movimentation test", func(t *testing.T) {
		account, transactions := createMovimentationData()

		movimentation := MovimentationFactory(account, transactions)

		operationDebitID := vo.GenerateID()
		operationDebitDescription := "Saque"
		operationDebitType := OperationDebit
		operationDebit := NewOperation(operationDebitID, operationDebitDescription, operationDebitType)

		transactionID := vo.GenerateID()
		amount, _ := vo.NewAmount(1.0)
		date := time.Now()
		transactionIntention := NewTransaction(transactionID, account, operationDebit, amount, date)

		check := movimentation.CheckMovimentation(transactionIntention)

		if check != nil {
			t.Errorf("Error when checking movimentation")
		}
	})
}
