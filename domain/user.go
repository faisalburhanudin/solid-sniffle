package domain

import "time"

type User struct {
	Id         int64
	Username   string
	Password   string
	Email      string
	ModifiedAt time.Time
	CreateAt   time.Time
}
