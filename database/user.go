package database

import (
	"database/sql"
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"log"
)

type UserDB struct {
	Db *sql.DB `inject:""`
}

func (udb UserDB) Get() []*domain.User {
	query := "SELECT id, username, email FROM users ORDER BY id DESC"
	rows, err := udb.Db.Query(query)
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
