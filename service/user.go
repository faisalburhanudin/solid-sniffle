package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type UserGetterByUsername interface {
	GetByUsername(username string) *domain.User
}

type UserDeleter interface {
	Delete(user domain.User)
}

type UserUpdater interface {
	Update(userId int, user domain.User) domain.User
}

// UserService used for inject interface to use in receiver
type UserService struct {
	UserGetterByUsername
	UserDeleter
	UserUpdater
}

type UserRepository interface {
	UserGetterByUsername
	UserDeleter
	UserUpdater
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository, repository, repository}
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
