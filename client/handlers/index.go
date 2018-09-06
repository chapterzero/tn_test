package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	log.Println(cwd)
	t := template.Must(template.ParseFiles(filepath.Join(cwd, "./tpl/index.gohtml")))
	t.Execute(w, nil)
}
