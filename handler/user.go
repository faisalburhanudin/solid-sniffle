package handler

import (
	"fmt"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"html"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

// Register handler
func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

	} else if r.Method == "GET" {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}

func (handler *UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	users := handler.UserService.Gets()
	for _, user := range users {
		fmt.Fprintf(w, "username: %s", user.Username)
	}
	return
}