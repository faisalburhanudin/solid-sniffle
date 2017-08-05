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
	Save(user *domain.User)
}

type UserGetter interface {
	Get(user domain.User) *domain.User
}

type UserService struct {
	usernameChecker UsernameChecker
	emailChecker    EmailChecker
	userSaver       UserSaver
	userGetter      UserGetter
}

var ErrorUsernameUsed = errors.New("Username sudah terpakai")
var ErrorEmailUsed = errors.New("Email sudah terpakai")
var ErrorUserNotFound = errors.New("User not found")

// Register new user
func (service UserService) Register(user *domain.User) error {
	// Check username used
	usernameUsed := service.usernameChecker.IsUsed(user.Username)
	if usernameUsed == true {
		return ErrorUsernameUsed
	}

	// Check email user
	emailUsed := service.emailChecker.IsUsed(user.Email)
	if emailUsed == true {
		return ErrorEmailUsed
	}

	// Save user
	service.userSaver.Save(user)

	return nil
}

// GetByUsername get user by username
func (service UserService) Get(userFilter domain.User) (*domain.User, error) {
	user := service.userGetter.Get(userFilter)
	if user == nil {
		return nil, ErrorUserNotFound
	}
	return user, nil
}
