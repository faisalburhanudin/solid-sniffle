package main

import (
	"net/http"
	"fmt"
	"html"
)

type User struct {
	id       int
	username string
}

type UserRepository interface {
	Save(user User)
}

type UserHandler struct {
	userRepository UserRepository
}

// register handler
func (handler *UserHandler) register(w http.ResponseWriter, r *http.Request) {
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

// register save user register
func register(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	email := r.FormValue("email")

	// username is mandatory
	if username == "" {
		http.Error(w, "Username harus di isi", http.StatusBadRequest)
		return
	}

	// email is mandatory
	if email == ""{
		http.Error(w, "Email harus di isi", http.StatusBadRequest)
		return
	}
}

// registerView render html form register
func registerView(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}