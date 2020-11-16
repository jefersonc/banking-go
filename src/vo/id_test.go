package vo

import (
	"fmt"
	"testing"
)

func TestGenerateIdSuccess(t *testing.T) {

	t.Run(fmt.Sprintf("ID test: generate a valid id"), func(t *testing.T) {
		id := GenerateID()

		_, err := NewID(id.Value())

		if err != nil {
			t.Errorf("Expected a valid id, got %v", err.Error())
		}
	})
}

func TestValidNewId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rawID string
	}{
		{
			rawID: "e5b320c1-9e58-4052-946a-f7bcfd5eb9cc",
		},
		{
			rawID: "0d4decdc-ef34-4379-bcff-6e217569b4e8",
		},
		{
			rawID: "f2ea50a1-e7cf-4a0d-8475-99179e5e5710",
		},
		{
			rawID: "f2c05d0d-2233-431a-8288-fbe07ab9f9ec",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ID test: %s", test.rawID), func(t *testing.T) {
			_, err := NewID(test.rawID)

			if err != nil {
				t.Errorf("Expected successfully ID generated, got %v", err.Error())
			}
		})
	}
}

func TestInvalidNewId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rawID string
	}{
		{
			rawID: "e5b320c1",
		},
		{
			rawID: "0d4decdc-ef34-bcff-6e217569b4e8",
		},
		{
			rawID: "f2ea50a1-e7cf-4a0d-8475-99179e5e5710-fea43",
		},
		{
			rawID: "123456789",
		},
		{
			rawID: "teste",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ID test: %s", test.rawID), func(t *testing.T) {
			_, err := NewID(test.rawID)

			if err == nil {
				t.Errorf("Expected failure on ID generation")
			}
		})
	}
}
