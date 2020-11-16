package src

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jefersonc/banking-go/src/adapters/repository"
	"github.com/jefersonc/banking-go/src/usecase"
)

// Application struct centralize application dependencies
type Application struct {
	router *mux.Router
}

// Bootstrap method init applicatiob with default dependencies
func Bootstrap(addr string) {
	a := &Application{}

	a.run(addr)
}

func (a Application) run(addr string) {
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
	//Defining middlewares
	// a.router.Use(middlewares.Logging)
	// a.router.Use(middlewares.ContentApplicationJson)

	accountRepository := repository.CreatePostgresAccountRepository()
	operationRepository := repository.CreatePostgresOperationRepository()
	transactionRepository := repository.CreatePostgresTransactionRepository()

	showAccount := usecase.NewFetchAccount(accountRepository)
	createAccount := usecase.NewCreateAccount(accountRepository)
	createTransaction := usecase.NewCreateTransaction(accountRepository, operationRepository, transactionRepository)

	//Registering routes
	a.router.HandleFunc("/accounts/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", showAccount.Invoke).Methods("GET")
	a.router.HandleFunc("/accounts", createAccount.Invoke).Methods("POST")
	a.router.HandleFunc("/accounts", createTransaction.Invoke).Methods("POST")
}
