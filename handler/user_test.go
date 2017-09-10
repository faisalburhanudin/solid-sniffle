package handler

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type UserRegistererMock struct {
	called bool
}

func (h UserRegistererMock) Register(user *domain.User) error {
	return nil
}

func TestUserHandler_Register(t *testing.T) {
	form := url.Values{}
	form.Add("username", "username")
	form.Add("email", "mail@host.com")
	form.Add("password", "pass")

	req := httptest.NewRequest(http.MethodPost, "http://localhost/Register", strings.NewReader(form.Encode()))
	w := httptest.NewRecorder()

	registerMock := UserRegistererMock{false}

	userHandler := UserHandler{
		TemplateDir:    "",
		UserRegisterer: registerMock,
	}
	userHandler.Register(w, req)
	if registerMock.called == false {
		t.Errorf("register mock not called")
	}
}

func TestUserHandler_Create_missing_param(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost/Register", nil)
	w := httptest.NewRecorder()

	userHandler := UserHandler{}
	userHandler.Register(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	type request struct {
		req http.ResponseWriter
		w   *http.Request
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("got: %v, want: %v.", resp.StatusCode, http.StatusBadRequest)
	}

	if string(body) != "Username harus di isi\n" {
		t.Errorf("got: %v, want: %v.", string(body), "Username harus di isi")
	}
}
