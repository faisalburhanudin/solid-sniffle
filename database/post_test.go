package database

import (
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestPost(t *testing.T) {
	db := NewSqliteMemory()

	CreatePostTable(db)

	// check zero result
	posts := GetListPost(db)
	if len(posts) != 0 {
		t.Errorf("got: %d, want: %d.", len(posts), 0)
	}

	// create new record
	link := "faisalburhanudin@github.com"
	title := "website"
	createdAt := time.Now()
	modifiedAt := time.Now()
	post := Post{
		link:       link,
		title:      title,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}
	AddPost(db, post)

	// get after one insert
	posts = GetListPost(db)
	if posts[0].link != link {
		t.Errorf("got: %v, want: %v.", posts[0].link, link)
	}
	if posts[0].title != title {
		t.Errorf("got: %v, want: %v.", posts[0].title, title)
	}
	if posts[0].createdAt != createdAt {
		t.Errorf("got: %v, want: %v.", posts[0].createdAt, createdAt)
	}
	if posts[0].modifiedAt != modifiedAt {
		t.Errorf("got: %v, want: %v.", posts[0].modifiedAt, modifiedAt)
	}
}
