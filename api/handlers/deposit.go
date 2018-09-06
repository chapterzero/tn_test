package handlers

import (
	"database/sql"
	"github.com/chapterzero/tn_test/api"
	"github.com/chapterzero/tn_test/api/codes"
	"github.com/chapterzero/tn_test/server"
	"net/http"
	"strconv"
)

func PostDepositHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := server.DbConnect()
	defer db.Close()
	postDepositAction(w, r, db)
}

func postDepositAction(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	pAccountNumber := r.FormValue("account_number")
	pAmount := r.FormValue("amount")

	// validate required parameter
	if pAccountNumber == "" || pAmount == "" {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "account number & amount are required",
		}
		api.WriteBadResponseError(w, errResponse)
	}

	// validate amount
	vAmount, err := strconv.Atoi(pAmount)
	if err != nil || vAmount <= 0 {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "amount should be integer and > 0",
		}
		api.WriteBadResponseError(w, errResponse)
	}
}
