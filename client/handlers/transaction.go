package handlers

import (
	"net/http"
	"text/template"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(GetTemplatePath("./tpl/transaction.gohtml")))
	t.Execute(w, nil)
}
