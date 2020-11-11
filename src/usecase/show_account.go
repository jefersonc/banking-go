package usecase

import (
	"net/http"

	"github.com/jefersonc/banking-go/src/ports/repository"
)

type ShowAccount struct {
	repository repository.AccountRepository
}

func (uc ShowAccount) Invoke(w http.ResponseWriter, r *http.Request) {
}

func NewShowAccount(repository repository.AccountRepository) {
	return ShowAccount{repository}
}
