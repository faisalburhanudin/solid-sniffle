package database

import (
	"database/sql"
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"log"
)

type UserAllGetter struct {
	Db *sql.DB `inject:""`
}

func (d *UserAllGetter) Get() []*domain.User {
	query := "SELECT id, username, email FROM users ORDER BY id DESC"
	rows, err := d.Db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// create slice with Post data
	var posts []*domain.User
	for rows.Next() {
		user := domain.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			log.Panic(err)
		}
		posts = append(posts, &user)
	}

	return posts
}

type UsernameChecker struct {
	Db *sql.DB `inject:""`
}

func (d UserAllGetter) IsUsed(username string) bool {
	query := "SELECT username FROM users WHERE username=?"
	rows, err := d.Db.Query(query, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

type EmailChecker struct {
	Db *sql.DB `inject:""`
}

func (d EmailChecker) IsUsed(email string) bool {
	query := "SELECT email FROM users WHERE email=?"
	rows, err := d.Db.Query(query, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

type UserSaver struct {
	Db *sql.DB `inject:""`
}

func (d UserSaver) Save(user *domain.User) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	res, err := d.Db.Exec(query, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	user.Id = lastInsertId
}
