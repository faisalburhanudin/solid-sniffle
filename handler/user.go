package handler

import (
	"fmt"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"html"
	"net/http"
)

type UserHandler struct {
	userRepository service.UserService
}

// Register handler
func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		register(w, r)
		return

	} else if r.Method == "GET" {
		registerView(w, r)
		return

	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}

// Register save user Register
func register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
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
}

// registerView render html form Register
func registerView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
