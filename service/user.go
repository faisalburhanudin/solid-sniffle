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

type UserGetterByUsername interface {
	GetByUsername(username string) *domain.User
}

type UserDeleter interface {
	Delete(user domain.User)
}

type UserUpdater interface {
	Update(userId int, user domain.User) domain.User
}

// error that indicate username already use
var ErrorUsernameUsed = errors.New("Username sudah terpakai")

// error that indicate email address already use
var ErrorEmailUsed = errors.New("Email sudah terpakai")

// UserService used for inject interface to use in receiver
type UserService struct {
	UsernameChecker      UsernameChecker `inject:""`
	EmailChecker         EmailChecker    `inject:""`
	UserSaver            UserSaver       `inject:""`
	UserGetterByUsername UserGetterByUsername
	UserDeleter          UserDeleter
	UserUpdater          UserUpdater
}

// Register new user
// this function will check if username already use or
// email address already use
func (s UserService) Register(user *domain.User) error {
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

// GetByUsername will retrieve data user by username
func (s UserService) GetByUsername(username string) domain.User {
	user := s.UserGetterByUsername.GetByUsername(username)
	return *user
}

// Delete user from application
func (s UserService) Delete(user domain.User) {
	s.UserDeleter.Delete(user)
}

// Update user detail
// userId parameter indicate id user will be update
// userPatch parameter will indicate which field will update so use fresh object
func (s UserService) Update(userId int, userPatch domain.User) domain.User {
	user := s.UserUpdater.Update(userId, userPatch)
	return user
}
