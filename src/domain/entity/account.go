package entity

// Account is entity/aggregator
type Account struct {
	id             int
	documentNumber string
}

// GetID is a getter for document_number attribute
func (a Account) GetID() int {
	return a.id
}

// GetDocumentNumber is a getter for document_number attribute
func (a Account) GetDocumentNumber() string {
	return a.documentNumber
}

// CreateAccount is a constructor
func CreateAccount(id int, documentNumber string) Account {
	return Account{id, documentNumber}
}
