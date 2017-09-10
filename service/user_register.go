package service

import (
	"errors"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type UsernameChecker interface {
	IsUsernameUsed(username string) bool
}

type EmailChecker interface {
	IsEmailUsed(email string) bool
}

type UserSaver interface {
	Save(user *domain.User)
}

// error that indicate username already use
var ErrorUsernameUsed = errors.New("Username sudah terpakai")

// error that indicate email address already use
var ErrorEmailUsed = errors.New("Email sudah terpakai")

type RegisterService struct {
	UsernameChecker
	EmailChecker
	UserSaver
}

type RegisterRepository interface {
	UsernameChecker
	EmailChecker
	UserSaver
}

func NewRegisterService(repository RegisterRepository) *RegisterService {
	return &RegisterService{repository, repository, repository}
}

// Register new user
// this function will check if username already use or
// email address already use
func (s RegisterService) Register(user *domain.User) error {
	// Check username used
	usernameUsed := s.UsernameChecker.IsUsernameUsed(user.Username)
	if usernameUsed == true {
		return ErrorUsernameUsed
	}

	// Check email user
	emailUsed := s.EmailChecker.IsEmailUsed(user.Email)
	if emailUsed == true {
		return ErrorEmailUsed
	}

	s.UserSaver.Save(user)
	return nil
}
