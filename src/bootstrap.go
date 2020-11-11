package src

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func (a *Application) registerHandlers() {
	//Defining middlewares
	a.router.Use(middlewares.Logging)
	a.router.Use(middlewares.ContentApplicationJson)

	accountRepository := adapters.CreatePostgresAccountRepository()

	showAccount := usecase.NewShowAccount(accountRepository)
	//Registering routes
	a.router.HandleFunc("/accounts/{id:[0-9]+}", showAccount.Invoke).Methods("GET")
}
