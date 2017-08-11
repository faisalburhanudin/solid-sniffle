package handler

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService `inject:""`
}

func (h UserHandler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Show all user
		h.Gets(w, r)
	case "POST":
		// Create new user
		h.Create(w, r)
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

// Create new user
// will return error if username
func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Email harus di isi", http.StatusBadRequest)
		return
	}

	// Save user
	user := domain.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	h.UserService.Register(&user)
	return
}

func (h UserHandler) Gets(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

}