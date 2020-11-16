package domain

import (
	"testing"

	"github.com/jefersonc/banking-go/src/vo"
)

func TestSuccessNewAccount(t *testing.T) {

	t.Run("Account instance test", func(t *testing.T) {
		id := vo.GenerateID()
		document, _ := vo.NewDocument("CPF", "42145390014")
		account := NewAccount(id, document)

		if account.GetID().Value() != id.Value() ||
			account.GetDocument().Number() != document.Number() ||
			account.GetDocument().Type() != document.Type() {
			t.Errorf("Error when checking account instance")
		}
	})
}
