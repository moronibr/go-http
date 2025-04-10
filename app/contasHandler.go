package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ContasHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/contasHandler.html"))
	tmpl.Execute(w, nil)
}

func ApiContasHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT name, email, password, registrationNumber, status FROM users")
		if err != nil {
			http.Error(w, "Erro ao consultar contas", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Name, &user.Email, &user.Password, &user.RegistrationNumber, &user.Status); err != nil {
				http.Error(w, "Erro ao ler contas", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func EditAccountHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/contas/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Erro ao decodificar", http.StatusBadRequest)
			return
		}

		query := `UPDATE users SET name=?, email=?, password=?, registration_number=?, status=? WHERE id=?`
		_, err = db.Exec(query, user.Name, user.Email, user.Password, user.RegistrationNumber, user.Status, id)
		if err != nil {
			http.Error(w, "Erro ao atualizar", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Conta atualizada com sucesso")
	}
}

func DeleteAccountHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/contas/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		query := `DELETE FROM users WHERE id=?`
		_, err = db.Exec(query, id)
		if err != nil {
			http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Conta deletada com sucesso")
	}
}
