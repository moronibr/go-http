package app

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ClientePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/cliente.html"))
	tmpl.Execute(w, nil)
}

type Cliente struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"data_nascimento"`
	Idade          int    `json:"idade"`
	Cidade         string `json:"cidade"`
	Estado         string `json:"estado"`
	Pais           string `json:"pais"`
	Ocupacao       string `json:"ocupacao"`
}

func AddClienteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var cliente Cliente
		err := json.NewDecoder(r.Body).Decode(&cliente)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			log.Println("Erro ao decodificar JSON:", err)
			return
		}

		log.Printf("Recebido: %+v\n", cliente)

		// Aqui você insere no banco. Certifique-se de que os campos batem com a tabela.
		_, err = db.Exec(`INSERT INTO cliente (nome, data_nascimento, idade, cidade, estado, pais, ocupacao)
						VALUES (?, ?, ?, ?, ?, ?, ?)`,
			cliente.Nome, cliente.DataNascimento, cliente.Idade, cliente.Cidade, cliente.Estado, cliente.Pais, cliente.Ocupacao)

		if err != nil {
			http.Error(w, "Erro ao inserir no banco de dados", http.StatusInternalServerError)
			log.Println("Erro ao inserir no banco:", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cliente adicionado com sucesso!"))
	}
}

func GetClientByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/api/clientes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var cliente Cliente
		err = db.QueryRow("SELECT id, nome, data_nascimento, idade, cidade, estado, pais, ocupacao FROM clientes WHERE id = ?", id).
			Scan(&cliente.ID, &cliente.Nome, &cliente.DataNascimento, &cliente.Idade, &cliente.Cidade, &cliente.Estado, &cliente.Pais, &cliente.Ocupacao)
		if err != nil {
			http.Error(w, "Cliente não encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cliente)
	}
}
