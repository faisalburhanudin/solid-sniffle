package database

import (
	"database/sql"
	"log"
	"time"
)

// Post structure data
type Post struct {
	id         int
	link       string
	title      string
	createdAt  time.Time
	modifiedAt time.Time
}

// GetListPost retrive post list in database
func GetListPost(db *sql.DB) []Post {
	// get from database
	sql := "SELECT link, title, createdAt, modifiedAt FROM posts ORDER BY id DESC"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// create slice with Post data
	var posts []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.link, &post.title, &post.createdAt, &post.modifiedAt); err != nil {
			log.Panic(err)
		}
		posts = append(posts, post)
	}

	return posts
}

// AddPost create new entry post in database
func AddPost(db *sql.DB, post Post) {
	sql := "INSERT INTO posts (id, link, title, createdAt, modifiedAt) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(sql, nil, post.link, post.title, post.createdAt, post.modifiedAt)
	if err != nil {
		log.Fatal(err)
	}
}

// CreatePostTable create schema database for post
func CreatePostTable(db *sql.DB) {
	sql := "CREATE TABLE posts ( " +
		"id int PRIMARY KEY, " +
		"link varchar(500), " +
		"title varchar(500), " +
		"createdAt DATETIME, " +
		"modifiedAt DATETIME);"
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}
