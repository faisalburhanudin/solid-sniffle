package main

import (
	"database/sql"
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/faisalburhanudin/solid-sniffle/database"
	"github.com/faisalburhanudin/solid-sniffle/handler"
	"github.com/faisalburhanudin/solid-sniffle/service"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	port := "8000"

	// Build DB
	db, err := sql.Open("mysql", "solid:pass@/solid")
	if err != nil {
		log.Error(err)
	}

	var g inject.Graph
	var userDB database.UserDB
	var userService service.UserService
	var userHandler handler.UserHandler

	// Inject singleton object
	err = g.Provide(
		&inject.Object{Value: &userDB},
		&inject.Object{Value: &userService},
		&inject.Object{Value: &userHandler},
		&inject.Object{Value: db},
	)
	if err != nil {
		log.Panic(err)
	}

	// Populate dependencies graph
	if err := g.Populate(); err != nil {
		log.Fatal(err)
	}

	// Register handler
	mux := http.NewServeMux()
	mux.HandleFunc("/register", userHandler.Register)
	mux.HandleFunc("/users", userHandler.ListUser)

	// Create middleware
	middle := negroni.New()
	middle.UseHandler(mux)
	// write request log
	middle.UseHandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{"method": r.Method, "endpoint": r.URL.Path}).Info("request")
	})

	// Build server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: middle,
	}

	// Running server
	log.WithFields(log.Fields{"port": port}).Info("HTTP Server running")
	log.Fatal(srv.ListenAndServe())
}
