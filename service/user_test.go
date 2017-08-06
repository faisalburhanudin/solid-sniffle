package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"reflect"
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
		UsernameChecker: mockUsernameChecker{isUsedReturn: true},
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
		UsernameChecker: mockUsernameChecker{isUsedReturn: false},
		EmailChecker:    mockEmailChecker{isUsedReturn: true},
	}
	err := userService.Register(&domain.User{})
	if err != ErrorEmailUsed {
		t.Errorf("got: %v, want: %v.", err, ErrorEmailUsed)
	}
}

type mockUserSaver struct {
	userId int64
}

func (mock mockUserSaver) Save(user *domain.User) {
	user.Id = mock.userId
}

func TestUserService_Register(t *testing.T) {
	userService := UserService{
		UsernameChecker: mockUsernameChecker{isUsedReturn: false},
		EmailChecker:    mockEmailChecker{isUsedReturn: false},
		UserSaver:       mockUserSaver{userId: 1},
	}
	user := domain.User{}
	userService.Register(&user)
	if user.Id != 1 {
		t.Errorf("got: %v, want: %v.", user.Id, 1)
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
		UserGetter: mockUserGetter{userReturn: nil},
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
		UserGetter: mockUserGetter{userReturn: &want},
	}
	got, _ := UserService.Get(domain.User{Username: "faisal"})
	if got != &want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

type mockUserAllGetter struct {
	usersReturn []*domain.User
}

func (mock mockUserAllGetter) Get() []*domain.User {
	return mock.usersReturn
}

func TestUserService_Gets(t *testing.T) {
	want := []*domain.User{
		{Username: "user1"},
		{Username: "user1"},
		{Username: "user1"},
	}
	UserService := UserService{
		UserAllGetter: mockUserAllGetter{},
	}
	got := UserService.Gets()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
