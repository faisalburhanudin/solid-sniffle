package database

import (
	"database/sql"
	"github.com/faisalburhanudin/solid-sniffle/domain"
	log "github.com/sirupsen/logrus"
)

type PostDb struct {
	Db *sql.DB
}

func NewPostDb(db *sql.DB) *PostDb {
	return &PostDb{Db: db}
}

// todo implement
func (d PostDb) Save(post *domain.Post) {

}

// todo implement
func (d PostDb) Delete(post domain.Post) {

}

func (d PostDb) GetPosts() []domain.Post {
	query := "SELECT id, text, user_id FROM posts ORDER BY id DESC"
	rows, err := d.Db.Query(query)
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()

	posts := []domain.Post{}
	for rows.Next() {
		post := domain.Post{}
		if err := rows.Scan(&post.Id, &post.Text, &post.UserId); err != nil {
			log.Error(err)
		}
	}

	return posts
}
