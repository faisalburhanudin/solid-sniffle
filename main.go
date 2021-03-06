package main

import (
	"database/sql"
	"github.com/faisalburhanudin/solid-sniffle/database"
	"github.com/faisalburhanudin/solid-sniffle/handler"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"github.com/faisalburhanudin/solid-sniffle/templates"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func httpLog(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		before := time.Now()
		h.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"method":   r.Method,
			"endpoint": r.URL.Path,
			"duration": time.Since(before),
		}).Info("request")
	})
}

func main() {
	addr := ":8000"

	// Build DB
	db, err := sql.Open("mysql", "solid:pass@/solid")
	if err != nil {
		log.Error(err)
	}

	// Check mysql connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot connect to mysql: %v", err)
	}

	postDb := database.NewPostDb(db)
	userDb := database.NewUserDb(db)

	postService := service.NewPostService(postDb)
	registerService := service.NewRegisterService(userDb)

	postHandler := handler.NewPostHandler(postService, templates.TemplateDir())
	userHandler := handler.NewUserHandler(registerService, templates.TemplateDir())

	// register routing
	var routing = []Routing{
		{"/register", userHandler.Register, []string{"GET", "POST"}},
		{"/", postHandler.List, []string{"GET"}},
		{"/user", userHandler.User, []string{"GET", "POST"}},
	}

	wrapper := []HttpWrapper{
		httpLog,
	}

	srv := NewServer(routing, wrapper)
	fs := http.FileServer(http.Dir("./static"))
	srv.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(addr, srv)
}
