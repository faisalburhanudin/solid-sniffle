package main

import (
	"database/sql"
	"github.com/facebookgo/inject"
	"github.com/faisalburhanudin/solid-sniffle/database"
	"github.com/faisalburhanudin/solid-sniffle/handler"
	"github.com/faisalburhanudin/solid-sniffle/service"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
	"net/http"
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
	port := 8000

	// Build DB
	db, err := sql.Open("mysql", "solid:pass@/solid")
	if err != nil {
		log.Error(err)
	}

	var g inject.Graph
	var userHandler handler.UserHandler

	// Inject singleton object
	err = g.Provide(
		&inject.Object{Value: &database.UserAllGetter{}},
		&inject.Object{Value: &database.UsernameChecker{}},
		&inject.Object{Value: &database.UserSaver{}},
		&inject.Object{Value: &service.UserService{}},
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

	// register routing
	routing := Routing{
		"/user":    userHandler.User,
	}

	wrapper := []HttpWrapper{
		httpLog,
	}

	srv := NewServer(port, routing, wrapper)
	srv.ListenAndServe()
}
