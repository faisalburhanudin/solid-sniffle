package handler

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"html/template"
	"net/http"
)

type UserRegisterer interface {
	Register(user *domain.User) error
}

type UserHandler struct {
	TemplateDir string
	UserRegisterer
}

func NewUserHandler(registerer UserRegisterer, templateDir string) *UserHandler {
	return &UserHandler{TemplateDir: templateDir, UserRegisterer: registerer}
}

// User root handler
func (h UserHandler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Show all user
		h.Gets(w, r)
	case "POST":
		// Register new user
		h.Register(w, r)
	case "PUT":
		// Update user
		h.Update(w, r)
	case "DELETE":
		// Delete user
		h.Delete(w, r)
	default:
		http.Error(w, "Method nt allowed", http.StatusMethodNotAllowed)
	}
}

// Register new user
// will return error if username
func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.registerGet(w, r)
	case "POST":
		h.registerPost(w, r)
	default:
		http.Error(w, "Method nt allowed", http.StatusMethodNotAllowed)
	}

	return
}

func (h UserHandler) registerGet(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		h.TemplateDir + "/front-base.html",
		h.TemplateDir + "/register.html",
	}
	tmpl := template.Must(template.ParseFiles(paths...))
	tmpl.Execute(w, nil)
}

// registerPost Handle register post
func (h UserHandler) registerPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	// username is mandatory
	if username == "" {
		http.Error(w, "Username harus di isi", http.StatusBadRequest)
		return
	}

	// email is mandatory
	if email == "" {
		http.Error(w, "Email harus di isi", http.StatusBadRequest)
		return
	}

	// Password is mandatory
	if password == "" {
		http.Error(w, "Password harus di isi", http.StatusBadRequest)
		return
	}

	// Save user
	user := domain.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	h.UserRegisterer.Register(&user)
}

func (h UserHandler) Gets(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
