package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type CreateTransfer struct {
	accountRepository     domain.AccountRepository
	operationRepository   domain.OperationRepository
	transactionRepository domain.TransactionRepository
}

type CreateTransferPayload struct {
	accountID       string
	operationTypeID string
	amount          float64
}

func (uc *CreateTransfer) Invoke(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	var payload CreateTransferPayload
	json.Unmarshal(body, &payload)

	accountId, _ := vo.NewID(payload.accountID)
	operationId, _ := vo.NewID(payload.operationTypeID)

	account, err := uc.accountRepository.Find(accountId)

	if err != nil {
		fmt.Println(err)
	}

	operation, err := uc.operationRepository.Find(operationId)

	if err != nil {
		fmt.Println(err)
	}

	amount, err := vo.NewAmount(payload.amount)

	if err != nil {
		fmt.Println(err)
	}

	transactionIntention := domain.NewTransaction(vo.GenerateID(), account, operation, amount, time.Now())
	transactions, err := uc.transactionRepository.FetchByAccount(account)

	if err != nil {
		fmt.Println(err)
	}

	movimentation := domain.MovimentationFactory(account, transactions)

	if movimentation.CheckMovimentation(transactionIntention, operation) != nil {
		fmt.Println(err)
	}

	if uc.transactionRepository.Push(transactionIntention) != nil {
		fmt.Println(err)
	}

	//success
	fmt.Println(account)
}

func NewCreateTransaction(account domain.AccountRepository, operation domain.OperationRepository, transaction domain.TransactionRepository) CreateTransfer {
	return CreateTransfer{account, operation, transaction}
}
