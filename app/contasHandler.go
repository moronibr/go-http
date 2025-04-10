package app

import (
	"html/template"
	"net/http"
)

func ContasHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/contasHandler.html"))
	tmpl.Execute(w, nil)
}
