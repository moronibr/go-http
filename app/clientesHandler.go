package app

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
)

func ClientesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/clientesHandler.html"))
	tmpl.Execute(w, nil)
}

func ApiClientesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nome, data_nascimento, idade, cidade, estado, pais, ocupacao FROM cliente")
		if err != nil {
			http.Error(w, "Erro ao consultar clientes", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var clientes []Cliente
		for rows.Next() {
			var cliente Cliente
			if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.DataNascimento, &cliente.Idade, &cliente.Cidade, &cliente.Estado, &cliente.Pais, &cliente.Ocupacao); err != nil {
				http.Error(w, "Erro ao ler clientes", http.StatusInternalServerError)
				return
			}
			clientes = append(clientes, cliente)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(clientes); err != nil {
			http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
			println("Erro no JSON:", err.Error()) // log aqui
			return
		}
	}
}
