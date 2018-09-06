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

type PostRegisterResponse struct {
	CustomerId    int64
	AccountNumber string
}

func PostRegisterHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := server.DbConnect()
	defer db.Close()
	postRegisterAction(w, r, db)
}

func postRegisterAction(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	amount := r.FormValue("amount")

	valid, vAmount := validatePostParam(firstName, lastName, amount, w)
	if !valid {
		return
	}

	dateString := time.Now().Format("2006-01-02")
	tx, _ := db.Begin()

	// create customer
	customerResult, err := tx.Exec(
		"INSERT INTO customer (first_name, last_name, join_date) VALUES(?,?,?)",
		firstName,
		lastName,
		dateString,
	)

	// handling if customer creation failed
	if err != nil {
		tx.Rollback()
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "Error occured when inserting to customer table: " + err.Error(),
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	// create account number
	accNumber := createNewAccountNumber(db)

	// get customer id from last inserted data
	customerId, _ := customerResult.LastInsertId()

	// create account
	_, err = tx.Exec(
		"INSERT INTO account (account_number, customer_id, opening_date, closing_date, balance) VALUES(?,?,?,?,?)",
		accNumber,
		customerId,
		dateString,
		nil,
		0,
	)

	// handling if account creation failed
	if err != nil {
		tx.Rollback()
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "Error occured when inserting to account table: " + err.Error(),
		}
		api.WriteBadResponseError(w, errResponse)
		return
	}

	// create transaction
	currentDate := time.Now()
	currentDate.Format(time.RFC3339)
	_, err = tx.Exec(
		"INSERT INTO transaction (account_number, description, amount, dtype, date) VALUES(?,?,?,?,?)",
		accNumber,
		"Opening balance",
		vAmount,
		"debit",
		currentDate,
	)

	// handling if account creation failed
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

	response := api.OkResponse{PostRegisterResponse{
		CustomerId:    customerId,
		AccountNumber: accNumber,
	}}

	api.WriteJsonResponse(w, response)
}

func validatePostParam(firstName, lastName, amount string, w http.ResponseWriter) (valid bool, vAmount int) {
	// validate required parameter
	if firstName == "" || lastName == "" || amount == "" {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "first name, last name & amount are required",
		}
		api.WriteBadResponseError(w, errResponse)
		return valid, vAmount
	}

	// validate amount
	vAmount, err := strconv.Atoi(amount)
	if err != nil || vAmount <= 0 {
		errResponse := api.ErrResponse{
			Code: codes.ErrValidation,
			Msg:  "amount should be integer and > 0",
		}
		api.WriteBadResponseError(w, errResponse)
		return valid, vAmount
	}

	valid = true
	return valid, vAmount
}

func createNewAccountNumber(db *sql.DB) (accountNumber string) {
	var dbAccNumber string
	db.QueryRow("SELECT max(account_number) as max FROM account").Scan(&dbAccNumber)

	if dbAccNumber == "" {
		dbAccNumber = "8001000"
	}

	numericAccNumber, _ := strconv.Atoi(dbAccNumber)
	numericAccNumber++
	accountNumber = strconv.Itoa(numericAccNumber)

	return accountNumber
}
