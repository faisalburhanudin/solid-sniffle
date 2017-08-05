package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"testing"
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
		emailChecker:    mockEmailChecker{isUsedReturn: true},
	}
	err := userService.Register(&domain.User{})
	if err != ErrorEmailUsed {
		t.Errorf("got: %v, want: %v.", err, ErrorEmailUsed)
	}
}

type mockUserGetter struct {
	userReturn *domain.User
}

func (mock mockUserGetter) Get(user domain.User) *domain.User {
	return mock.userReturn
}

func TestUserService_Get_NotFound(t *testing.T) {
	UserService := UserService{
		userGetter: mockUserGetter{userReturn: nil},
	}
	_, err := UserService.Get(domain.User{Username: "faisal"})
	if err != ErrorUserNotFound {
		t.Errorf("got: %v, want: %v.", err, ErrorUserNotFound)
	}
}

func TestUserService_Get(t *testing.T) {
	want := domain.User{
		Username: "faisal",
	}
	UserService := UserService{
		userGetter: mockUserGetter{userReturn: &want},
	}
	got, _ := UserService.Get(domain.User{Username: "faisal"})
	if got != &want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
