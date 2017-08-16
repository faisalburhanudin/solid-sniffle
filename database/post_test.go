// +build integration
package database

import (
	"testing"
	"flag"
	"database/sql"
	"reflect"
	"github.com/faisalburhanudin/solid-sniffle/domain"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
)

var dsn = flag.String("dsn", "solid:pass@tcp(127.0.0.1:3307)/solidtest", "database source name for testing")

func TestPostDb_GetPosts(t *testing.T) {
	db := openDB(t)
	createSchema(t, db)

	postDb := PostDb{Db: db}
	posts := postDb.GetPosts()

	want := []domain.Post{}
	assertEqual(t, posts, want)

	dropSchema(t, db)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("got: %v, want: %v.", a, b)
	}
}

func openDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		t.Error(err)
	}
	return db
}

func createSchema(t *testing.T, db *sql.DB) {
	file, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		t.Error(err)
	}

	var requests = []string{}

	for _, query := range strings.Split(string(file), ";") {
		clearQuery := strings.Replace(query, "\n", "", -1)
		if clearQuery != "" {
			requests = append(requests, clearQuery)
		}

	}

	for _, request := range requests {
		if _, err := db.Exec(request); err != nil {
			t.Error(err)
		}
	}
}

func dropSchema(t *testing.T, db *sql.DB) {
	if _, err := db.Exec("DROP TABLE users, posts"); err != nil {
		t.Error(err)
	}
}
