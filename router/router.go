package router

import (
	aHandler "github.com/chapterzero/tn_test/api/handlers"
	cHandler "github.com/chapterzero/tn_test/client/handlers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	registerRoutes(r)
	return r
}

func registerRoutes(r *mux.Router) {
	// client routes
	r.HandleFunc("/", cHandler.IndexHandler).Name("Index").Methods("GET")
	r.HandleFunc("/deposit", cHandler.DepositHandler).Name("Deposit").Methods("GET")
	r.HandleFunc("/register", cHandler.RegisterHandler).Name("Register").Methods("GET")
	r.HandleFunc("/transaction", cHandler.TransactionHandler).Name("Transaction").Methods("GET")

	// api routes
	r.HandleFunc("/api/register", aHandler.PostRegisterHandler).Name("PostRegister").Methods("POST")
	r.HandleFunc("/api/deposit", aHandler.PostDepositHandler).Name("PostDeposit").Methods("POST")
}
