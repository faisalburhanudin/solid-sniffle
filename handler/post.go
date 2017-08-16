package handler

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"html/template"
	"net/http"
)

type PostGetter interface {
	GetPosts() []domain.Post
}

type PostHandler struct {
	PostGetter PostGetter `inject:""`
}

func NewPostHandler(postGetter PostGetter) *PostHandler {
	return &PostHandler{PostGetter: postGetter}
}

// List write post data to html response writer
func (h PostHandler) List(w http.ResponseWriter, r *http.Request) {
	posts := h.PostGetter.GetPosts()
	paths := []string{
		"templates/front-base.tmpl",
		"templates/post-list.tmpl",
	}
	tmpl := template.Must(template.ParseFiles(paths...))
	tmpl.Execute(w, posts)
}
