package api

import (
	"testing"
)

func TestErrResponseErrFunc(t *testing.T) {
	err := ErrResponse {
		Code: 400,
		Path: "/api/post/customer",
		Msg: "Customer name is required",
	}

	expected := "Error when processing your request: Customer name is required"
	actual := err.Error()
	if actual != expected {
		t.Errorf("Expected \"%v\", got \"%v\"", expected, actual)
	}
}
