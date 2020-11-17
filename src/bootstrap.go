package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jefersonc/banking-go/src/adapters/repository"
	"github.com/jefersonc/banking-go/src/domain"
	"github.com/jefersonc/banking-go/src/infra/database"
	"github.com/jefersonc/banking-go/src/usecase"
)

// Application struct centralize application dependencies
type Application struct {
	router                *mux.Router
	accountRepository     domain.AccountRepository
	operationRepository   domain.OperationRepository
	transactionRepository domain.TransactionRepository
}

// Bootstrap method init applicatiob with default dependencies
func Bootstrap(addr string) {
	client := database.NewPostgresClient()

	accountRepository := repository.CreatePostgresAccountRepository(client)
	operationRepository := repository.CreatePostgresOperationRepository(client)
	transactionRepository := repository.CreatePostgresTransactionRepository(client)

	a := &Application{
		mux.NewRouter(),
		accountRepository,
		operationRepository,
		transactionRepository,
	}

	a.run(addr)
}

func (a *Application) Responder(res http.ResponseWriter, payload interface{}, err error) {
	res.Header().Add("Content-Type", "application/json")

	if err != nil {
		switch err.(type) {
		case *usecase.UserError:
			res.WriteHeader(http.StatusBadRequest)
		case *usecase.DomainError:
			res.WriteHeader(http.StatusUnprocessableEntity)
		case *usecase.ApplicationError:
			res.WriteHeader(http.StatusInternalServerError)
		default:
			res.WriteHeader(http.StatusInternalServerError)
		}

		res.Write([]byte(err.Error()))
	}

	encoded, _ := json.Marshal(payload)

	res.Write(encoded)
}

func (a *Application) run(addr string) {
	a.router = mux.NewRouter()

	a.registerHandlers()

	log.Println("Listening to port " + addr)

	server := &http.Server{
		Handler:      a.router,
		Addr:         addr,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func (a *Application) registerHandlers() {
	a.router.HandleFunc("/accounts/{id}", a.fetchAccountHandler).Methods("GET")
	a.router.HandleFunc("/accounts", a.createAccountHandler).Methods("POST")
	a.router.HandleFunc("/transactions", a.createTransactionHandler).Methods("POST")
}

func (a *Application) createAccountHandler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	var payload usecase.CreateAccountPayload
	json.Unmarshal(body, &payload)

	uc := usecase.NewCreateAccount(a.accountRepository)

	response, err := uc.Execute(payload)

	a.Responder(res, response, err)
}

func (a *Application) fetchAccountHandler(res http.ResponseWriter, req *http.Request) {
	uc := usecase.NewFetchAccount(a.accountRepository)
	queryId := mux.Vars(req)["id"]

	payload := usecase.FetchAccountPayload{ID: queryId}
	response, err := uc.Execute(payload)

	a.Responder(res, response, err)
}

func (a *Application) createTransactionHandler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	var payload usecase.CreateTransactionPayload
	json.Unmarshal(body, &payload)

	uc := usecase.NewCreateTransaction(a.accountRepository, a.operationRepository, a.transactionRepository)

	response, err := uc.Execute(payload)

	a.Responder(res, response, err)
}
