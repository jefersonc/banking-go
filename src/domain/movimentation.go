package domain

import "errors"

type Movimentation struct {
	account      *Account
	transactions *[]Transaction
}

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

func (m *Movimentation) CheckMovimentation(transactionIntention *Transaction, operation *Operation) error {
	if operation.GetFinality() == OperationCredit {
		return nil
	}

	balance := m.GetBalance()

	if balance < transactionIntention.GetAmount().Value() {
		return ErrInsufficientFunds
	}

	return nil
}

func (m *Movimentation) GetBalance() float64 {
	var balance float64 = 0.0

	for _, transaction := range *m.transactions {
		if transaction.operation.GetFinality() == OperationCredit {
			balance += transaction.amount.Value()
			continue
		}

		balance -= transaction.amount.Value()
	}

	return balance
}

func MovimentationFactory(account *Account, transactions *[]Transaction) *Movimentation {
	return &Movimentation{account, transactions}
}
