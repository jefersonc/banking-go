package entity

// OperationType is entity/aggregator
type OperationType struct {
	id          int
	description string
}

// GetID is a getter for id attribute
func (o OperationType) GetID() int {
	return o.id
}

// GetDescription is a getter for description attribute
func (o OperationType) GetDescription() string {
	return o.description
}

// CreateOperationType is a constructor
func CreateOperationType(id int, documentNumber string) OperationType {
	return OperationType{id, documentNumber}
}
