package service

import (
	"errors"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type UsernameChecker interface {
	IsUsed(username string) bool
}

type EmailChecker interface {
	IsUsed(email string) bool
}

type UserSaver interface {
	Save(user *domain.User) *domain.User
}

type UserService struct {
	usernameChecker UsernameChecker
	emailChecker    EmailChecker
	userSaver       UserSaver
}

var ErrorUsernameUsed = errors.New("Username sudah terpakai")
var ErrorEmailUsed = errors.New("Email sudah terpakai")

// Register new user
func (service UserService) Register(user *domain.User) (error) {
	// Check username used
	UsernameUsed := service.usernameChecker.IsUsed(user.Username)
	if UsernameUsed == true {
		return ErrorUsernameUsed
	}

	// Check email user
	EmailUsed := service.emailChecker.IsUsed(user.Email)
	if EmailUsed == true {
		return ErrorEmailUsed
	}

	// Save user
	service.userSaver.Save(user)

	return nil
}
