package session

import (
	"errors"
	"net/http"
)

// SessionData representa os dados da sessão
type SessionData struct {
	Name     string
	Password string
}

var sessions = map[string]SessionData{}

// SetSession cria uma nova sessão e a armazena no cookie
func SetSession(w http.ResponseWriter, name, password string) {
	cookie := &http.Cookie{
		Name:     "session_user",
		Value:    name,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Armazena os dados da sessão no mapa
	sessions[name] = SessionData{
		Name:     name,
		Password: password,
	}
}

// GetSession recupera os dados da sessão a partir do cookie
func GetSession(r *http.Request) (SessionData, error) {
	cookie, err := r.Cookie("session_user")
	if err != nil {
		return SessionData{}, err
	}

	session, exists := sessions[cookie.Value]
	if !exists {
		return SessionData{}, errors.New("sessão não encontrada")
	}

	return session, nil
}

// DeleteSession remove a sessão do mapa e limpa o cookie
func DeleteSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session_user",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // Expira o cookie
	}
	http.SetCookie(w, cookie)

	// Remove os dados da sessão
	delete(sessions, cookie.Value)
}
