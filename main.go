package main

import (
	"fmt"
	"go-http/app"
	"go-http/db"
	"log"
	"net/http"
)

func main() {
	// Conectar ao banco de dados
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbConn.Close()

	// Servir arquivos estáticos (CSS, JS, imagens, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Configurar as rotas
	http.HandleFunc("/", app.IndexHandler)        // Rota para a página index
	http.HandleFunc("/login", app.LoginHandler)   // Rota para login
	http.HandleFunc("/logout", app.LogoutHandler) // Rota para logout

	// Iniciar o servidor
	fmt.Println("Servidor rodando em http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
