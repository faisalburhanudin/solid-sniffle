package testing

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

var SQL_PATH string

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	current := path.Dir(filename)
	SQL_PATH = path.Join(current, "../database/schema.sql")
}

// CreateDB create schema, will return db connection pointer and
// function tear down
func CreateDB(t *testing.T) (*sql.DB, func()) {
	dsn := os.Getenv("DSN_TEST")

	if dsn == "" {
		t.Skip("no environment variable DSN_TEST for database testing: skip test")
	}
	// Register new connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}

	// open schema file
	file, err := ioutil.ReadFile(SQL_PATH)
	if err != nil {
		t.Fatal(err)
	}

	// create sql query based on file
	var requests = []string{}
	for _, query := range strings.Split(string(file), ";") {
		clearQuery := strings.Replace(query, "\n", "", -1)
		if clearQuery != "" {
			requests = append(requests, clearQuery)
		}

	}

	// execution query
	for _, request := range requests {
		if _, err := db.Exec(request); err != nil {
			t.Fatal(err)
		}
	}

	// delete all tables
	tearDown := func() {

		// get tables
		rows, err := db.Query("SHOW TABLES")
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()
		tables := []string{}
		for rows.Next() {
			table := ""
			if err := rows.Scan(&table); err != nil {
				t.Fatal(table)
			}
			tables = append(tables, table)
		}

		query := "DROP TABLE IF EXISTS " + strings.Join(tables, ",")
		if _, err := db.Exec(query); err != nil {
			t.Fatal(err)
		}
		db.Close()
	}

	return db, tearDown
}
