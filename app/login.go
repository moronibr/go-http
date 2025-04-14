package app

import (
	"go-http/session" // Importa o pacote de sessão
	"html/template"
	"net/http"
	// Para comparar a senha hash
)

// LoginPageHandler exibe o formulário de login
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/login.html")
	if err != nil {
		http.Error(w, "Erro ao carregar a página de login", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// LoginHandler valida as credenciais e cria a sessão
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Obtém os dados do formulário de login
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Aqui você deve validar o nome de usuário e senha com o banco de dados
	// Suponhamos que o banco de dados tenha um usuário "admin" com a senha "admin"
	// Em produção, você usaria um banco de dados para fazer a verificação real
	if username == "admin" && password == "admin" {
		// Cria a sessão para o usuário
		session.SetSession(w, username, password)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	// Se as credenciais forem inválidas
	http.Error(w, "Nome de usuário ou senha inválidos", http.StatusUnauthorized)
}
