package handlers

import (
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

	postDepositAction(rr, req, db)

	if rr.Code != 400 {
		t.Errorf("Expected 400 error code, got %v", rr.Code)
	}
}
