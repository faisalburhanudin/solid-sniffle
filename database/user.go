package database

import (
	"database/sql"
	"github.com/faisalburhanudin/solid-sniffle/domain"
	log "github.com/sirupsen/logrus"
)

type UserDb struct {
	Db *sql.DB
}

// NewUserDb user database handler
func NewUserDb(db *sql.DB) *UserDb {
	return &UserDb{db}
}

// Get user list
func (d *UserDb) Get() []*domain.User {
	query := "SELECT id, username, email FROM users ORDER BY id DESC"
	rows, err := d.Db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// create slice with user data
	var users []*domain.User
	for rows.Next() {
		user := domain.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			log.Panic(err)
		}
		users = append(users, &user)
	}

	return users
}

// IsUsernameUsed check if username already use
func (d UserDb) IsUsernameUsed(username string) bool {
	query := "SELECT username FROM users WHERE username=?"
	rows, err := d.Db.Query(query, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

// IsEmailUsed check if email already use
func (d UserDb) IsEmailUsed(email string) bool {
	query := "SELECT email FROM users WHERE email=?"
	rows, err := d.Db.Query(query, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

// Save new user
func (d UserDb) Save(user *domain.User) {
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

// Get user by username
func (d UserDb) GetByUsername(username string) *domain.User {
	query := "SELECT id, username, email FROM users ORDER BY id DESC"
	rows, err := d.Db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := domain.User{}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			log.Error(err)
		}
	}

	return &user
}
