package vo

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidIDPattern = errors.New("The ID format is invalid.")

type ID struct {
	value string
}

func GenerateID() *ID {
	id := uuid.New()
	return &ID{id.String()}
}

func NewID(id string) (*ID, error) {
	_, err := uuid.Parse(id)

	if err == nil {
		return &ID{id}, nil
	}

	return nil, ErrInvalidIDPattern
}

func (id *ID) Value() string {
	return id.value
}
