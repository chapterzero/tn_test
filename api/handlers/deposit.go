package handlers

import (
	"database/sql"
	"github.com/chapterzero/tn_test/api"
	"github.com/chapterzero/tn_test/api/codes"
	"github.com/chapterzero/tn_test/server"
	"net/http"
	"strconv"
	"time"
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
		return
	}

	// validate amount
	vAmount, err := strconv.Atoi(pAmount)
	if err != nil || vAmount <= 0 {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "amount should be integer and > 0",
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	// validate account_number by checking database
	var dbAccNumber string
	db.QueryRow("SELECT account_number FROM account where account_number = ?", pAccountNumber).Scan(&dbAccNumber)

	if dbAccNumber == "" {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "invalid account_number",
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	currentDate := time.Now()
	currentDate.Format(time.RFC3339)

	tx, _ := db.Begin()
	result, err := tx.Exec(
		"INSERT INTO transaction (account_number, description, amount, dtype, date) VALUES(?,?,?,?,?)",
		dbAccNumber,
		"Cash deposit",
		vAmount,
		"debit",
		currentDate,
	)

	if err != nil {
		tx.Rollback()
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "Error occured when inserting to transaction table: " + err.Error(),
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	tx.Commit()

	transactionId, _ := result.LastInsertId()
	response := api.OkResponse{
		Data: transactionId,
	}

	api.WriteJsonResponse(w, response)
}
