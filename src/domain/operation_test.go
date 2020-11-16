package domain

import (
	"testing"

	"github.com/jefersonc/banking-go/src/vo"
)

func TestSuccessNewOperation(t *testing.T) {

	t.Run("Operation instance test", func(t *testing.T) {
		id := vo.GenerateID()
		description := "Depósito"
		operationType := OperationCredit
		operation := NewOperation(id, description, operationType)

		if operation.GetID().Value() != id.Value() ||
			operation.GetDescription() != description ||
			operation.GetFinality() != operationType {
			t.Errorf("Error when checking operation instance")
		}
	})
}
