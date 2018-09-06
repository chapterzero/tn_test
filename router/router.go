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
	r.HandleFunc("/", cHandler.IndexHandler).Name("Index").Methods("GET")
	r.HandleFunc("/api/customer", aHandler.PostCustomerHandler).Name("PostCustomer").Methods("POST")
}
