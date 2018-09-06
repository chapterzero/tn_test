package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestTransactionHandler(t *testing.T) {
	os.Chdir("../..")

	req, _ := http.NewRequest("GET", "/transaction", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TransactionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong http code, expected %v, got %v", http.StatusOK, status)
	}

	os.Chdir("client/handlers")
}
