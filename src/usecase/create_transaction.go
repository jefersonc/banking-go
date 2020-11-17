package usecase

import (
	"time"

	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/vo"
)

type (
	// CreateTransaction is a usecase definition
	CreateTransaction struct {
		accountRepository     domain.AccountRepository
		operationRepository   domain.OperationRepository
		transactionRepository domain.TransactionRepository
	}

	// CreateTransactionPayload is a usecase parameters definition
	CreateTransactionPayload struct {
		AccountID       string  `json:"account_id"`
		OperationTypeID string  `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}

	// CreateTransactionResponse is a usecase response definition
	CreateTransactionResponse struct {
		TransactionID string  `json:"transaction_id"`
		AccountID     string  `json:"account_id"`
		OperationID   string  `json:"operation_id"`
		Date          string  `json:"date"`
		Amount        float64 `json:"amount"`
		Finality      string  `json:"finality"`
	}
)

func (uc *CreateTransaction) Execute(payload CreateTransactionPayload) (*CreateTransactionResponse, error) {
	amount, err := vo.NewAmount(payload.Amount)

	if err != nil {
		return nil, NewUserError(err)
	}

	accountID, err := vo.NewID(payload.AccountID)

	if err != nil {
		return nil, NewUserError(err)
	}

	operationID, err := vo.NewID(payload.OperationTypeID)

	if err != nil {
		return nil, NewUserError(err)
	}

	account, err := uc.accountRepository.Find(accountID)

	if err != nil {
		return nil, NewApplicationError(err)
	}

	operation, err := uc.operationRepository.Find(operationID)

	if err != nil {
		return nil, NewApplicationError(err)
	}

	transactionIntention := domain.NewTransaction(vo.GenerateID(), account, operation, amount, time.Now())
	transactions, err := uc.transactionRepository.FetchByAccount(account)

	if err != nil {
		return nil, NewApplicationError(err)
	}

	movimentation := domain.MovimentationFactory(account, transactions)

	err = movimentation.CheckMovimentation(transactionIntention)

	if err != nil {
		return nil, NewDomainError(err)
	}

	err = uc.transactionRepository.Push(transactionIntention)

	if err != nil {
		return nil, NewApplicationError(err)
	}

	return uc.output(transactionIntention), nil
}

func (uc *CreateTransaction) output(transaction *domain.Transaction) *CreateTransactionResponse {
	return &CreateTransactionResponse{
		TransactionID: transaction.GetID().Value(),
		AccountID:     transaction.GetAccount().GetID().Value(),
		OperationID:   transaction.GetOperation().GetID().Value(),
		Date:          transaction.GetDate().Format("2006-01-02T15:04:05-0700"),
		Amount:        transaction.GetAmount().Value(),
		Finality:      transaction.GetOperation().GetFinality(),
	}
}

// NewCreateTransaction is a constructor
func NewCreateTransaction(account domain.AccountRepository, operation domain.OperationRepository, transaction domain.TransactionRepository) CreateTransaction {
	return CreateTransaction{account, operation, transaction}
}
