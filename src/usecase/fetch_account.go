package usecase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type FetchAccount struct {
	repository domain.AccountRepository
}

func (uc FetchAccount) Invoke(w http.ResponseWriter, r *http.Request) {
	requestId := string(mux.Vars(r)["id"])

	id, err := vo.NewID(requestId)

	if err != nil {
		fmt.Println(err)
	}

	account, err := uc.repository.Find(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account)
}

func NewFetchAccount(repository domain.AccountRepository) FetchAccount {
	return FetchAccount{repository}
}
