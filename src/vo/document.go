package vo

import (
	"errors"

	"github.com/klassmann/cpfcnpj"
)

var (
	ErrDocumentTypeInvalid  = errors.New("The document type is invalid (valid values: CPF and CNPJ).")
	ErrDocumentValueInvalid = errors.New("The document value is invalid.")
)

type Document struct {
	documentType string
	number       string
}

func NewDocument(documentType string, number string) (*Document, error) {

	if documentType != "CPF" && documentType != "CNPJ" {
		return nil, ErrDocumentTypeInvalid
	}

	validated := false

	if documentType == "CPF" {
		validated = cpfcnpj.ValidateCPF(number)
	}

	if documentType == "CNPJ" {
		validated = cpfcnpj.ValidateCNPJ(number)
	}

	if validated {
		return &Document{documentType: documentType, number: number}, nil
	}

	return nil, ErrDocumentValueInvalid
}

func (d Document) Type() string {
	return d.documentType
}
func (d Document) Number() string {
	return d.number
}
