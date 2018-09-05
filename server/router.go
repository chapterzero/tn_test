package server

import (
	"github.com/chapterzero/tn_test/api/handlers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customer", handlers.PostCustomerHandler).Name("PostCustomer").Methods("POST")
	r.HandleFunc("/api/customer_type", handlers.GetCustomerTypeHandler).Name("GetCustomerType").Methods("GET")

	return r
}
