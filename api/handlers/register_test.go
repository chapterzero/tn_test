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

func TestPostRegisterActionWithoutParameter(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/register", nil)
	rr := httptest.NewRecorder()
	db, _, _ := sqlmock.New()

	defer db.Close()

	postRegisterAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostRegisterActionInvalidAmount(t *testing.T) {
	data := url.Values{}
	data.Set("first_name", "Yusuf")
	data.Set("last_name", "Irwandi")
	data.Add("amount", "invalid numb")
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, _, _ := sqlmock.New()

	defer db.Close()

	postRegisterAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestGenerateNewAccountNumberDbExist(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	column := []string{"account_number"}
	mock.ExpectQuery("SELECT (.+) FROM account").
		WillReturnRows(sqlmock.NewRows(column).AddRow("1001"))

	actual := createNewAccountNumber(db)

	if actual != "1002" {
		t.Errorf("Expected 1002, got %v", actual)
	}
}

func TestGenerateNewAccountNumberDbNotExist(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM account").
		WillReturnRows(sqlmock.NewRows([]string{}))

	actual := createNewAccountNumber(db)

	if actual != "8001001" {
		t.Errorf("Expected 8001001, got %v", actual)
	}
}

func TestPostRegisterShouldRollbackIfInsertCustomerError(t *testing.T) {
	data := url.Values{}
	data.Set("first_name", "Yusuf")
	data.Set("last_name", "Irwandi")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO customer (.+)").
		WillReturnError(errors.New("insert to database error"))
	mock.ExpectRollback()

	postRegisterAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostRegisterShouldRollbackIfInsertAccountError(t *testing.T) {
	data := url.Values{}
	data.Set("first_name", "Yusuf")
	data.Set("last_name", "Irwandi")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO customer (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO account (.+)").
		WillReturnError(errors.New("insert to account error"))
	mock.ExpectQuery("SELECT max\\(account_number\\) FROM account").
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectRollback()

	postRegisterAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostRegisterShouldRollbackIfInsertTransactionError(t *testing.T) {
	data := url.Values{}
	data.Set("first_name", "Yusuf")
	data.Set("last_name", "Irwandi")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO customer (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO account (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO transaction (.+)").
		WillReturnError(errors.New("insert to transaction error"))
	mock.ExpectQuery("SELECT max\\(account_number\\) FROM account").
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectRollback()

	postRegisterAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}

func TestPostRegisterCommit(t *testing.T) {
	data := url.Values{}
	data.Set("first_name", "Yusuf")
	data.Set("last_name", "Irwandi")
	data.Add("amount", "5000")
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	db, mock, _ := sqlmock.New()

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO customer (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO account (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO transaction (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT max\\(account_number\\) FROM account").
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectCommit()

	postRegisterAction(rr, req, db)

	if rr.Code != 200 {
		t.Errorf("Expected 200 status code, got %v", rr.Code)
	}
}
