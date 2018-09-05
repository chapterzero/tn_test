package handlers

import (
	"github.com/chapterzero/tn_test/api"
	"github.com/chapterzero/tn_test/api/codes"
	"github.com/chapterzero/tn_test/server"
	"log"
	"net/http"
)

func GetCustomerTypeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := server.DbConnect()
	if err != nil {
		errResponse := api.ErrResponse{
			Code: codes.ErrDbCannotConnect,
			Msg:  "Database error: " + err.Error(),
		}
		api.WriteInteralServerError(w, errResponse)
	}

	log.Println("Customer type handler", server.GetDbConfig())
	// db,_ := sql.Open("mysql", "root:some@tcp(127.0.0.1:33061)/tn_test")
	// err := db.Ping()
	// if err != nil {
	// 	payload := api.ErrResponse{1000, "/path", err.Error()}
	// 	api.WriteJsonResponse(w, api.ServerErrorResponse{payload})
	// }
}
