package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	os.Chdir("../..")

	req, _ := http.NewRequest("GET", "/register", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong http code, expected %v, got %v", http.StatusOK, status)
	}

	os.Chdir("client/handlers")
}
