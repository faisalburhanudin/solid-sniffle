package main

import (
	"testing"
)

type mockUsernameChecker struct {
	isUsedReturn bool
}

func (check mockUsernameChecker) IsUsed(username string) bool {
	return check.isUsedReturn
}

func TestUserService_RegisterUsernameUsed(t *testing.T) {
	service := UserApi{
		usernameChecker: mockUsernameChecker{isUsedReturn: true},
	}
	_, err := service.Register(User{})
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
	service := UserApi{
		usernameChecker: mockUsernameChecker{isUsedReturn: false},
		emailChecker: mockEmailChecker{isUsedReturn: true},
	}
	_, err := service.Register(User{})
	if err != ErrorEmailUsed {
		t.Errorf("got: %v, want: %v.", err, ErrorEmailUsed)
	}
}
