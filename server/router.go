package server

import (
	"github.com/chapterzero/tn_test/api/handlers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PostCustomerHandler)
	r.HandleFunc("/test", handlers.GetCustomerTypeHandler)

	return r
}
