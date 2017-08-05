package main

import (
	"fmt"
	"github.com/faisalburhanudin/solid-sniffle/handler"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"github.com/faisalburhanudin/solid-sniffle/database"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "8000"
	// Build DB
	db, err := sql.Open("mysql", "solid:pass@/solid")
	if err != nil {
		log.Error(err)
	}

	userDB := database.UserDB{
		Db: db,
	}

	// Build service
	userService := service.UserService{
		UserAllGetter: userDB,
	}

	// Create handler
	userHandler := &handler.UserHandler{UserService: &userService}

	// Register handler
	mux := http.NewServeMux()
	mux.HandleFunc("/register", userHandler.Register)
	mux.HandleFunc("/users", userHandler.ListUser)

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
