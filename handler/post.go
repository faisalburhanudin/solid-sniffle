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
	TemplateDir string
	PostGetter  PostGetter
}

func NewPostHandler(postGetter PostGetter, templateDir string) *PostHandler {
	return &PostHandler{PostGetter: postGetter, TemplateDir: templateDir}
}

// List write post data to html response writer
func (h PostHandler) List(w http.ResponseWriter, r *http.Request) {
	posts := h.PostGetter.GetPosts()
	paths := []string{
		h.TemplateDir + "/front-base.tmpl",
		h.TemplateDir + "/post-list.tmpl",
	}
	tmpl := template.Must(template.ParseFiles(paths...))
	tmpl.Execute(w, posts)
}
