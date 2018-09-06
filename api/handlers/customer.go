package handlers

import (
	"github.com/chapterzero/tn_test/api"
	"net/http"
)

func PostCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	payload := map[string]string{
		"id":   "1",
		"name": "yusuf",
	}
	api.WriteJsonResponse(w, api.OkResponse{payload})
}
