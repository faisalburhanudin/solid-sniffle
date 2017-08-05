package service

import (
	"testing"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type mockUsernameChecker struct {
	isUsedReturn bool
}

func (check mockUsernameChecker) IsUsed(username string) bool {
	return check.isUsedReturn
}

func TestUserService_RegisterUsernameUsed(t *testing.T) {
	userService := UserService{
		usernameChecker: mockUsernameChecker{isUsedReturn: true},
	}
	err := userService.Register(&domain.User{})
	if err != ErrorUsernameUsed {
		t.Errorf("got: %v, want: %v.", err, ErrorUsernameUsed)
	}
}

type mockEmailChecker struct {
	isUsedReturn bool
}

func (check mockEmailChecker) IsUsed(email string) bool {
	return check.isUsedReturn
}

func TestUserService_RegisterEmailUser(t *testing.T) {
	userService := UserService{
		usernameChecker: mockUsernameChecker{isUsedReturn: false},
		emailChecker: mockEmailChecker{isUsedReturn: true},
	}
	err := userService.Register(&domain.User{})
	if err != ErrorEmailUsed {
		t.Errorf("got: %v, want: %v.", err, ErrorEmailUsed)
	}
}
