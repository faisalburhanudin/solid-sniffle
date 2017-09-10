package database

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	test "github.com/faisalburhanudin/solid-sniffle/testing"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
)

func TestPostDb_GetPosts(t *testing.T) {
	db, tearDown := test.CreateDB(t)
	defer tearDown()

	postDb := PostDb{Db: db}
	posts := postDb.GetPosts()

	want := []domain.Post{}
	if reflect.DeepEqual(posts, want) != true {
		t.Errorf("got %v: want %v", posts, want)
	}
}
