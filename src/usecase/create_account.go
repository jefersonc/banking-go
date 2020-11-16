package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type CreateAccount struct {
	repository domain.AccountRepository
}

type CreateAccountPayload struct {
	documentNumber string
}

func (uc *CreateAccount) Invoke(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	var payload CreateAccountPayload
	json.Unmarshal(body, &payload)

	fmt.Printf(string(body))

	document, _ := vo.NewDocument("CPF", payload.documentNumber)

	account := domain.NewAccount(vo.GenerateID(), document)

	err := uc.repository.Push(account)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account)
}

func NewCreateAccount(repository domain.AccountRepository) CreateAccount {
	return CreateAccount{repository}
}
