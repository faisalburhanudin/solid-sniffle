package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
)

func TestUserHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost/register", nil)
	w := httptest.NewRecorder()

	userHandler := UserHandler{}
	userHandler.register(w, req)

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