package api

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJsonResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	someResponse := OkResponse{
		Data: "value",
	}

	WriteJsonResponse(rr, someResponse)

	actualBody := strings.TrimSpace(rr.Body.String())
	expectedBody := "{\"data\":\"value\"}"
	expectedContentType := "application/json"

	if actualBody != expectedBody {
		t.Errorf("Expected body %v, got '%v'", expectedBody, actualBody)
	}

	if rr.Header().Get("content-type") != expectedContentType {
		t.Errorf("Expected header content type %v, got %v", expectedContentType, rr.Header().Get("content-type"))
	}
}

func TestInternalServerErrorResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	someErrResponse := ErrResponse{
		Code: 1,
		Msg:  "Some Error Occured",
	}

	WriteInteralServerError(rr, someErrResponse)
	actualBody := strings.TrimSpace(rr.Body.String())
	actualHeader := rr.Header().Get("content-type")

	expectedBody := "{\"code\":1,\"msg\":\"Some Error Occured\"}"

	if actualBody != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, actualBody)
	}

	if actualHeader != "application/json" {
		t.Errorf("Expected header content type %v, got %v", "application/json", actualHeader)
	}

	if rr.Code != 500 {
		t.Errorf("Expected status code %v, got %v", 500, rr.Code)
	}
}
