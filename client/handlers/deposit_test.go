package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDepositHandler(t *testing.T) {
	os.Chdir("../..")

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DepositHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong http code, expected %v, got %v", http.StatusOK, status)
	}

	os.Chdir("client/handlers")
}
