package app

import (
	"html/template"
	"net/http"
)

func ClientesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/clientesHandler.html"))
	tmpl.Execute(w, nil)
}
