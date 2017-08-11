package main

import (
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// signature for function wrapper
type HttpWrapper func(h http.HandlerFunc) http.HandlerFunc

// signature for endpoint routing
// ex: "/user":funcHandler
type Routing map[string]http.HandlerFunc

// A Server defines parameters for running an HTTP server.
type Server struct {
	port   int
	mux    *http.ServeMux
}

// NewServer allocates and return Server
// this will assign logger, port,
// register routing who pointed endpoint and function handler,
// wrapper handler fox example using for middleware
func NewServer(port int, routing Routing, wrapper []HttpWrapper) *Server {
	server := &Server{
		mux:    http.NewServeMux(),
		port:   port,
	}

	// Register handler function and endpoint to mux
	for url, handlerFunc := range routing {
		// for each handler wrap with wrapper provide
		for _, w := range wrapper {
			handlerFunc = w(handlerFunc)
		}

		// finally register to mux
		server.mux.HandleFunc(url, handlerFunc)
	}

	return server
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
	log.Fatal(srv.ListenAndServe())
}
