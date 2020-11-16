package vo

import (
	"fmt"
	"testing"
)

func TestNewDocumentSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		documentType string
		number       string
		result       *Document
	}{
		{
			documentType: "CPF",
			number:       "01714126137",
			result:       &Document{"CPF", "01714126137"},
		},
		{
			documentType: "CPF",
			number:       "08411504948",
			result:       &Document{"CPF", "08411504948"},
		},
		{
			documentType: "CNPJ",
			number:       "34278073000161",
			result:       &Document{"CNPJ", "34278073000161"},
		},
		{
			documentType: "CNPJ",
			number:       "08843964000103",
			result:       &Document{"CNPJ", "08843964000103"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Document test: %s, %s", test.documentType, test.number), func(t *testing.T) {
			document, _ := NewDocument(test.documentType, test.number)

			if test.result.Type() != document.Type() ||
				test.result.Number() != document.Number() {
				t.Errorf("Expected %v, got %v", test.result, document)
			}
		})
	}
}

func TestNewDocumentError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		documentType string
		number       string
		err          error
	}{
		{
			documentType: "RG",
			number:       "12345678901",
			err:          ErrDocumentTypeInvalid,
		},
		{
			documentType: "cpf",
			number:       "12345678901",
			err:          ErrDocumentTypeInvalid,
		},
		{
			documentType: "cnpj",
			number:       "12345678901",
			err:          ErrDocumentTypeInvalid,
		},
		{
			documentType: "CNH",
			number:       "12345678901",
			err:          ErrDocumentTypeInvalid,
		},
		{
			documentType: "CPF",
			number:       "12345678901",
			err:          ErrDocumentValueInvalid,
		},
		{
			documentType: "CPF",
			number:       "123456789",
			err:          ErrDocumentValueInvalid,
		},
		{
			documentType: "CPF",
			number:       "123456789",
			err:          ErrDocumentValueInvalid,
		},
		{
			documentType: "CPF",
			number:       "896.688.868-28",
			err:          ErrDocumentValueInvalid,
		},
		{
			documentType: "CNPJ",
			number:       "89668886828",
			err:          ErrDocumentValueInvalid,
		},
		{
			documentType: "CNPJ",
			number:       "34.199.719/0001-15",
			err:          ErrDocumentValueInvalid,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Document test: %s, %s", test.documentType, test.number), func(t *testing.T) {
			_, err := NewDocument(test.documentType, test.number)

			if test.err.Error() != err.Error() {
				t.Errorf("Expected %v, got %v", test.err.Error(), err.Error())
			}
		})
	}
}
