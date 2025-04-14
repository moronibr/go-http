package middleware

import (
	"net/http"
)

// AuthMiddleware vai verificar o cookie de sessão
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_user")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// Se estiver autenticado, passa pro handler real
		next.ServeHTTP(w, r)
	})
}
