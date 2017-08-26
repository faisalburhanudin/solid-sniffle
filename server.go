package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// signature for function wrapper
type HttpWrapper func(h http.HandlerFunc) http.HandlerFunc

// signature for endpoint routing
// ex: "/user":funcHandler
type Routing struct {
	Endpoint string
	Handler  http.HandlerFunc
	Method   []string
}

// NewServer allocates and return Server
// this will assign logger, port,
// register routing who pointed endpoint and function handler,
// wrapper handler fox example using for middleware
func NewServer(routing []Routing, wrapper []HttpWrapper) *mux.Router {
	m := mux.NewRouter()

	// Register handler function and endpoint to mux
	for _, route := range routing {
		// for each handler wrap with wrapper provide
		for _, w := range wrapper {
			route.Handler = w(route.Handler)
		}

		// finally register to mux
		m.HandleFunc(route.Endpoint, route.Handler).Methods(route.Method...)
	}

	return m
}
