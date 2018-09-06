package api

import (
	"encoding/json"
	"net/http"
)

func WriteBadResponseError(w http.ResponseWriter, errResponse ErrResponse) {
	w.WriteHeader(http.StatusBadRequest)
	WriteJsonResponse(w, errResponse)
}

func WriteInteralServerError(w http.ResponseWriter, errResponse ErrResponse) {
	w.WriteHeader(http.StatusInternalServerError)
	WriteJsonResponse(w, errResponse)
}

func WriteJsonResponse(w http.ResponseWriter, response Response) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
