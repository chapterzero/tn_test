package api

import (
	"net/http"
	"encoding/json"
)

type Response interface{}

type OkResponse struct {
	Data  interface{}  `json:"data"`
}

type BadResponse struct {
	Err ErrResponse    `json:"err"`
}

type ServerErrorResponse struct {
	Err ErrResponse    `json:"err"`
}

type ErrResponse struct {
	Code   int         `json:"code"`
	Path   string      `json:"path"`
	Msg    string      `json:"msg"`
}

func (e *ErrResponse) Error() (string) {
	return "Error when processing your request: " + e.Msg
}

func WriteJsonResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
