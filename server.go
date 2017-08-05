package main

import (
	"fmt"
	"github.com/faisalburhanudin/solid-sniffle/handler"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	port := "8000"

	// Create handler
	userHandler := handler.UserHandler{}

	// Register handler
	mux := http.NewServeMux()
	mux.HandleFunc("/register", userHandler.Register)

	// Create middleware
	middle := negroni.New()
	middle.UseHandlerFunc(LogRequest)
	middle.UseHandler(mux)

	// Build server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: middle,
	}

	// Running server
	log.WithFields(log.Fields{"port": port}).Info("HTTP Server running")
	log.Fatal(srv.ListenAndServe())
}

func LogRequest(_ http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{"method": r.Method, "endpoint": r.URL.Path}).Info("request")
}
