package main

import (
	"net/http"
)

func main() {
	// Create handler
	userHandler := UserHandler{}

	// Wiring register
	mux := http.NewServeMux()
	mux.HandleFunc("/register", userHandler.register)

	// Serve http
	http.ListenAndServe(":8000", mux)
}

