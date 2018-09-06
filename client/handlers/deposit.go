package handlers

import (
	"net/http"
	"text/template"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(GetTemplatePath("./tpl/deposit.gohtml")))
	t.Execute(w, nil)
}
