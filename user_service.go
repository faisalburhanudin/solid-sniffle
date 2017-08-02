package main

import "errors"

type UsernameChecker interface {
	IsUsed(username string) bool
}

type EmailChecker interface {
	IsUsed(email string) bool
}

type UserSaver interface {
	Save(user User) User
}

type UserApi struct {
	usernameChecker UsernameChecker
	emailChecker    EmailChecker
	userSaver       UserSaver
}

var ErrorUsernameUsed = errors.New("Username sudah terpakai")
var ErrorEmailUsed = errors.New("Email sudah terpakai")

// Register new user
func (service UserApi) Register(user User) (User, error) {
	// Check username used
	UsernameUsed := service.usernameChecker.IsUsed(user.username)
	if UsernameUsed == true {
		return User{}, ErrorUsernameUsed
	}

	// Check email user
	EmailUsed := service.emailChecker.IsUsed(user.email)
	if EmailUsed == true {
		return User{}, ErrorEmailUsed
	}

	// Save user
	service.userSaver.Save(user)

	return user, nil
}
