package app

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
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
