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

func TestPostDepositActionWithoutParameter(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/deposit", nil)
	rr := httptest.NewRecorder()
	db, _, _ := sqlmock.New()

	defer db.Close()

	postDepositAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostDepositActionInvalidAmount(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "acc number")
	data.Add("amount", "invalid numb")
	req, _ := http.NewRequest("POST", "/api/deposit", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, _, _ := sqlmock.New()

	defer db.Close()

	postDepositAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostDepositActionInvalidAccountNumber(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "accnumber")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/deposit", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	postDepositAction(rr, req, db)

	mock.ExpectQuery("SELECT (.+) FROM account where account_number (.+)").
		WithArgs("accnumber").
		WillReturnRows(sqlmock.NewRows([]string{}))

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostDepositActionValidAccountRollbackIfInsertError(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "accnumber")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/deposit", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	columns := []string{"account_number"}
	mock.ExpectQuery("SELECT account_number FROM account where account_number(.+)").
		WithArgs("accnumber").
		WillReturnRows(sqlmock.NewRows(columns).AddRow("accnumber"))
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO transaction (.+)").
		WillReturnError(errors.New("insert to database error"))
	mock.ExpectRollback()

	postDepositAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostDepositActionSuccess(t *testing.T) {
	data := url.Values{}
	data.Set("account_number", "accnumber")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/deposit", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	columns := []string{"account_number"}
	mock.ExpectQuery("SELECT account_number FROM account where account_number(.+)").
		WithArgs("accnumber").
		WillReturnRows(sqlmock.NewRows(columns).AddRow("accnumber"))
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO transaction (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	postDepositAction(rr, req, db)

	if rr.Code != 200 {
		t.Errorf("Expected 200 error code, got %v", rr.Code)
	}

}
