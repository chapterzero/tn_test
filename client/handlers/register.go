package handlers

import (
	"net/http"
	"text/template"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(GetTemplatePath("./tpl/register.gohtml")))
	t.Execute(w, nil)
}
