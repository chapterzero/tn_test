package handlers

import (
	"errors"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetTransactionActionWithoutParameter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/transaction", nil)
	rr := httptest.NewRecorder()
	db, _, _ := sqlmock.New()

	defer db.Close()

	getTransactionAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestGetTransactionActionAccountQueryError(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "accnumber")
	req, _ := http.NewRequest("POST", "/api/transaction", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()

	db, mock, _ := sqlmock.New()

	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM account WHERE (.+)").
		WithArgs("accnumber").
		WillReturnError(errors.New("error while select to database"))

	getTransactionAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestGetTransactionActionQueryReturnError(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "accnumber")
	req, _ := http.NewRequest("POST", "/api/transaction", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()

	db, mock, _ := sqlmock.New()

	defer db.Close()

	column := []string{"account_number", "amount"}
	mock.ExpectQuery("SELECT (.+) FROM account WHERE (.+)").
		WithArgs("accnumber").
		WillReturnRows(sqlmock.NewRows(column).AddRow("1001", 50000))
	mock.ExpectQuery("SELECT (.+) FROM transaction WHERE (.+)").
		WithArgs("accnumber").
		WillReturnError(errors.New("error while select to database"))

	getTransactionAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestGetTransactionActionQuery(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "accnumber")
	req, _ := http.NewRequest("POST", "/api/transaction", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()

	db, mock, _ := sqlmock.New()

	defer db.Close()

	column := []string{"account_number", "amount"}
	mock.ExpectQuery("SELECT (.+) FROM account WHERE (.+)").
		WithArgs("accnumber").
		WillReturnRows(sqlmock.NewRows(column).AddRow("1001", 50000))
	column = []string{"account_number", "description", "amount", "dtype", "date"}
	mock.ExpectQuery("SELECT (.+) FROM transaction WHERE (.+)").
		WithArgs("accnumber").
		WillReturnRows(sqlmock.NewRows(column).AddRow("1001", "Cash", 50000, "debit", "2018-08-08 16:30:00"))

	getTransactionAction(rr, req, db)

	if rr.Code != 200 {
		t.Errorf("Expected 200 http code, got %v", rr.Code)
	}
}
