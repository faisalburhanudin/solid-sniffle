package handler

import (
	"github.com/faisalburhanudin/solid-sniffle/service"
	"html/template"
	"net/http"
)

type PostHandler struct {
	PostService *service.PostService `inject:""`
}

// List display post
func (h PostHandler) List(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"templates/front-base.tmpl",
		"templates/post-list.tmpl",
	}
	tmpl := template.Must(template.ParseFiles(paths...))
	tmpl.Execute(w, nil)
}
