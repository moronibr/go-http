package app

import (
	"go-http/session" // Substitua pelo caminho correto
	"html/template"
	"net/http"
)

// IndexHandler exibe a página inicial, se o usuário estiver autenticado
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	sessionData, err := session.GetSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Redireciona para login se não houver sessão
		return
	}

	// Exemplo: se o nome do usuário for admin, renderiza a página de admin
	if sessionData.Name != "admin" {
		http.Error(w, "Acesso não autorizado", http.StatusForbidden)
		return
	}

	// Renderiza a página index.html
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
