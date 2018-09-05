package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"testing"
)

var r *mux.Router

func init() {
	r = CreateRouter()
}

func TestCreateRouterPostCustomer(t *testing.T) {
	postCustomerRoute := r.Get("PostCustomer")
	method, _ := postCustomerRoute.GetMethods()
	pathExp, _ := postCustomerRoute.GetPathRegexp()

	assertSingleMethod(method, "POST", t)
	assertPath(pathExp, "/api/customer", t)
}

func TestCreateRouterGetCustomerType(t *testing.T) {
	getCustomerType := r.Get("GetCustomerType")
	method, _ := getCustomerType.GetMethods()
	pathExp, _ := getCustomerType.GetPathRegexp()

	assertSingleMethod(method, "GET", t)
	assertPath(pathExp, "/api/customer_type", t)
}

func assertSingleMethod(actual []string, expected string, t *testing.T) {
	if len(actual) != 1 || actual[0] != expected {
		t.Errorf("Route should only accept %v, got %v", expected, actual)
	}
}

func assertPath(actual string, expected string, t *testing.T) {
	expected = fmt.Sprintf("^%v$", expected)
	if actual != expected {
		t.Errorf("Wrong path, expected %v, got %v", expected, actual)
	}
}
