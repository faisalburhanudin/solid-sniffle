package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"testing"
)

type mockUsernameChecker struct {
	isUsedReturn bool
}

func (check mockUsernameChecker) IsUsernameUsed(username string) bool {
	return check.isUsedReturn
}

func TestUserService_RegisterUsernameUsed(t *testing.T) {
	userService := RegisterService{
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

func (check mockEmailChecker) IsEmailUsed(email string) bool {
	return check.isUsedReturn
}

func TestUserService_RegisterEmailUser(t *testing.T) {
	userService := RegisterService{
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
	userService := RegisterService{
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

func (mock mockUserGetter) GetByUsername(username string) *domain.User {
	return mock.userReturn
}

func TestUserService_Get(t *testing.T) {
	want := domain.User{
		Username: "faisal",
	}
	UserService := UserService{
		UserGetterByUsername: mockUserGetter{userReturn: &want},
	}
	got := UserService.GetByUsername("faisal")
	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

type mockUserDeleter struct {
	isCalled bool
}

func (mock *mockUserDeleter) Delete(user domain.User) {
	mock.isCalled = true
}

func TestUserService_Delete(t *testing.T) {
	mock := mockUserDeleter{}
	userService := UserService{
		UserDeleter: &mock,
	}
	userService.Delete(domain.User{})
	if mock.isCalled != true {
		t.Errorf("got: %v, want: %v.", mock.isCalled, true)
	}
}

type mockUserUpdater struct {
	userReturn domain.User
}

func (mock mockUserUpdater) Update(userId int, user domain.User) domain.User {
	return mock.userReturn
}

func TestUserService_Update(t *testing.T) {
	userReturn := domain.User{}

	mock := mockUserUpdater{userReturn: userReturn}
	userService := UserService{
		UserUpdater: mock,
	}
	user := userService.Update(1, domain.User{})
	if user != userReturn {
		t.Errorf("got: %v, want: %v.", user, userReturn)
	}
}
