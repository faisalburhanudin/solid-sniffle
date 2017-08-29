package handler

import (
	"testing"
	test "github.com/faisalburhanudin/solid-sniffle/testing"
	"github.com/faisalburhanudin/solid-sniffle/database"
	"github.com/faisalburhanudin/solid-sniffle/service"
	"net/http/httptest"
	"database/sql"
	"github.com/faisalburhanudin/solid-sniffle/templates"
)

func postHandlerFactory(db *sql.DB) *PostHandler {
	postDb := database.NewPostDb(db)
	postService := service.NewPostService(postDb)
	return NewPostHandler(postService, templates.TemplateDir())
}

func TestPostHandler_List(t *testing.T) {
	db, tearDown := test.CreateDB(t)
	defer tearDown()

	postHandler := postHandlerFactory(db)

	req := httptest.NewRequest("GET", "http://1.2.3.4/", nil)
	w := httptest.NewRecorder()

	postHandler.List(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("got %v, want %v", resp.StatusCode, 200)
	}
}
