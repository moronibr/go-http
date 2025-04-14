package app

import (
	"go-http/session" // Substitua pelo caminho correto
	"net/http"
)

// LogoutHandler lida com o logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Remove a sessão
	session.DeleteSession(w)

	// Redireciona para a página de login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
