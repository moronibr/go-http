package app

import (
	"go-http/db"
	"html/template"
	"log"
	"net/http"
)

func ContaHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/conta.html"))
	tmpl.Execute(w, nil)
}

type User struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	RegistrationNumber string `json:"registrationNumber"`
	Status             string `json:"status"`
}

func AddAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	users := User{
		Name:               r.FormValue("name"),
		Email:              r.FormValue("email"),
		Password:           r.FormValue("password"),
		RegistrationNumber: r.FormValue("registrationNumber"),
		Status:             r.FormValue("status"),
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		log.Println("Erro ao conectar ao banco:", err)
		return
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO users (name, email, password, registrationNumber, status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Erro ao preparar query", http.StatusInternalServerError)
		log.Println("Erro ao preparar query:", err)
		return
	}
	defer stmt.Close()

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	registrationNumber := r.FormValue("registrationNumber")
	statusStr := r.FormValue("status")

	// Validação
	if name == "" || email == "" || password == "" || registrationNumber == "" || statusStr == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	_, err = stmt.Exec(users.Name, users.Email, users.Password, users.RegistrationNumber, users.Status)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco", http.StatusInternalServerError)
		log.Println("Erro ao inserir no banco:", err) // <-- AQUI!!!
		return
	}

	w.Write([]byte("Conta adicionada com sucesso!"))
}
