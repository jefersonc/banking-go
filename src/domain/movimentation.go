package domain

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Movimentation struct {
	account      *Account
	transactions []*Transaction
}

func (m *Movimentation) GetAccount() *Account {
	return m.account
}

func (m *Movimentation) GetTransactions() []*Transaction {
	return m.transactions
}

func (m *Movimentation) CheckMovimentation(transactionIntention *Transaction) error {
	if transactionIntention.GetOperation().GetFinality() == OperationCredit {
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

	for _, transaction := range m.transactions {
		if transaction.operation.GetFinality() == OperationCredit {
			balance += transaction.amount.Value()
			continue
		}

		balance -= transaction.amount.Value()
	}

	return balance
}

func MovimentationFactory(account *Account, transactions []*Transaction) *Movimentation {
	return &Movimentation{account, transactions}
}
