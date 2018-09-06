package handlers

import (
	"database/sql"
	"github.com/chapterzero/tn_test/api"
	"github.com/chapterzero/tn_test/api/codes"
	"github.com/chapterzero/tn_test/server"
	"log"
	"net/http"
)

type TransactionResponse struct {
	AccountNumber string
	Balance       sql.NullString
	Transactions  []Transaction
}

type Transaction struct {
	Id            string
	AccountNumber string
	Description   string
	Amount        sql.NullString
	Dtype         string
	Date          string
}

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := server.DbConnect()
	defer db.Close()
	getTransactionAction(w, r, db)
}

func getTransactionAction(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	accountNumber := r.FormValue("account_number")

	if accountNumber == "" {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "account number is required",
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	var trResponse TransactionResponse

	// get account balance
	err := db.QueryRow("SELECT account_number, balance FROM account WHERE account_number = ?", accountNumber).Scan(&trResponse.AccountNumber, &trResponse.Balance)
	if err != nil {
		errResponse := api.ErrResponse{
			Code: codes.ErrDbQuery,
			Msg:  err.Error(),
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	rows, err := db.Query("SELECT * FROM transaction WHERE account_number = ?", accountNumber)
	if err != nil {
		errResponse := api.ErrResponse{
			Code: codes.ErrDbQuery,
			Msg:  err.Error(),
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var trRow Transaction
		err = rows.Scan(&trRow.Id, &trRow.AccountNumber, &trRow.Description, &trRow.Amount, &trRow.Dtype, &trRow.Date)
		if err != nil {
			log.Println(err)
		}
		trResponse.Transactions = append(trResponse.Transactions, trRow)
	}

	okResponse := api.OkResponse{trResponse}
	api.WriteJsonResponse(w, okResponse)
}
