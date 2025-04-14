package main

import (
	"fmt"
	"log"
	"net/http"

	"go-http/app"
	"go-http/db"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbConn.Close()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.LoginPageHandler)

	http.HandleFunc("/login", app.LoginHandler(dbConn))

	http.HandleFunc("/logout", app.LogoutHandler)

	http.HandleFunc("/index", app.IndexHandler)

	fmt.Println("Servidor rodando em http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
