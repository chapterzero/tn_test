package handlers

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(GetTemplatePath("./tpl/index.gohtml")))
	t.Execute(w, nil)
}
