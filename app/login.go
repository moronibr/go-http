package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/login.html")
	if err != nil {
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		password := r.FormValue("password")

		var storedHash string
		err := db.QueryRow("SELECT password FROM users WHERE name = ?", name).Scan(&storedHash)
		if err != nil {
			fmt.Println("Erro ao buscar no banco:", err)
			http.Error(w, "Dados Inválidos", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
		if err != nil {
			fmt.Println("Senha incorreta!")
			http.Error(w, "Dados Inválidos", http.StatusUnauthorized)
			return
		}

		// Criar cookie de sessão
		cookie := &http.Cookie{
			Name:  "session_user",
			Value: name,
			Path:  "/",
			// Opcional: HttpOnly, Secure, SameSite
		}
		http.SetCookie(w, cookie)

		fmt.Println("Login bem-sucedido! Redirecionando para /index")
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
}
