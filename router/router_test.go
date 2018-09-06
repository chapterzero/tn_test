package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"testing"
)

var r *mux.Router

func init() {
	r = CreateRouter()
}

func TestCreateRouterPostDeposit(t *testing.T) {
	route := r.Get("PostDeposit")
	method, _ := route.GetMethods()
	pathExp, _ := route.GetPathRegexp()

	assertSingleMethod(method, "POST", t)
	assertPath(pathExp, "/api/deposit", t)
}

func TestCreateRouterPostRegister(t *testing.T) {
	route := r.Get("PostRegister")
	method, _ := route.GetMethods()
	pathExp, _ := route.GetPathRegexp()

	assertSingleMethod(method, "POST", t)
	assertPath(pathExp, "/api/register", t)
}

func TestCreateRouterIndex(t *testing.T) {
	route := r.Get("Index")
	method, _ := route.GetMethods()
	pathExp, _ := route.GetPathRegexp()

	assertSingleMethod(method, "GET", t)
	assertPath(pathExp, "/", t)
}

func TestCreateRouterDeposit(t *testing.T) {
	route := r.Get("Deposit")
	method, _ := route.GetMethods()
	pathExp, _ := route.GetPathRegexp()

	assertSingleMethod(method, "GET", t)
	assertPath(pathExp, "/deposit", t)
}

func TestCreateRouterRegister(t *testing.T) {
	route := r.Get("Register")
	method, _ := route.GetMethods()
	pathExp, _ := route.GetPathRegexp()

	assertSingleMethod(method, "GET", t)
	assertPath(pathExp, "/register", t)
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
