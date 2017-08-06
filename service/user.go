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

type UserAllGetter interface {
	Get() []*domain.User
}

var ErrorUsernameUsed = errors.New("Username sudah terpakai")
var ErrorEmailUsed = errors.New("Email sudah terpakai")
var ErrorUserNotFound = errors.New("User not found")

type UserService struct {
	UsernameChecker UsernameChecker `inject:""`
	EmailChecker    EmailChecker    `inject:""`
	UserSaver       UserSaver       `inject:""`
	UserGetter      UserGetter
	UserAllGetter   UserAllGetter `inject:""`
}

// Register new user
func (service *UserService) Register(user *domain.User) error {
	// Check username used
	usernameUsed := service.UsernameChecker.IsUsed(user.Username)
	if usernameUsed == true {
		return ErrorUsernameUsed
	}

	// Check email user
	emailUsed := service.EmailChecker.IsUsed(user.Email)
	if emailUsed == true {
		return ErrorEmailUsed
	}

	// Save user
	service.UserSaver.Save(user)

	return nil
}

// Get single user get user by username
func (service *UserService) Get(userFilter domain.User) (*domain.User, error) {
	user := service.UserGetter.Get(userFilter)
	if user == nil {
		return nil, ErrorUserNotFound
	}
	return user, nil
}

// Gets multiple users
func (service *UserService) Gets() []*domain.User {
	return service.UserAllGetter.Get()
}
