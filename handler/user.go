package handler

import (
	"fmt"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"net/http"
	"html/template"
	"log"
)

type UserHandler struct {
	UserService *service.UserService `inject:""`
}

type RegisterPage struct {
	Action string
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
		t, err := template.ParseFiles("template/register.html")
		if err != nil {
			log.Panic(err)
		}

		page := RegisterPage{Action: "/register"}
		t.Execute(w, &page)
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
