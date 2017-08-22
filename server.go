package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

// A Server defines parameters for running an HTTP server.
type Server struct {
	port int
	mux  *mux.Router
}

// NewServer allocates and return Server
// this will assign logger, port,
// register routing who pointed endpoint and function handler,
// wrapper handler fox example using for middleware
func NewServer(port int, routing []Routing, wrapper []HttpWrapper) *Server {
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

	return &Server{
		mux:  m,
		port: port,
	}
}

// ListenAndServe listens on the TCP network address srv.Addr
func (s *Server) ListenAndServe() {
	// format address with port
	addr := fmt.Sprintf("localhost:%d", s.port)

	// Create server handler
	srv := http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	// listen port and serve
	log.Infof("listen on %v", addr)
	log.Fatal(srv.ListenAndServe())
}
