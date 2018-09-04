package handlers

import (
	"github.com/chapterzero/tn_test/api"
	"net/http"
)

func GetCustomerTypeHandler(w http.ResponseWriter, r*http.Request) {
	db,_ := sql.Open("mysql", "root:some@tcp(127.0.0.1:33061)/tn_test")
	err := db.Ping()
	if err != nil {
		payload := api.ErrResponse{1000, "/path", err.Error()}
		api.WriteJsonResponse(w, api.ServerErrorResponse{payload})
	}
}
