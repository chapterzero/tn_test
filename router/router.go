package router

import (
	"github.com/gorilla/mux"
	"github.com/chapterzero/tn_test/api/handlers"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PostCustomerHandler)
	r.HandleFunc("/test", handlers.GetCustomerTypeHandler)

	return r
}
