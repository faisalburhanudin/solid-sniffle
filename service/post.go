package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type PostSaver interface {
	Save(post *domain.Post)
}

type PostDeleter interface {
	Delete(post domain.Post)
}

// PostService using for inject dependencies via interface
type PostService struct {
	PostSaver   PostSaver
	PostDeleter PostDeleter
}

// Create new post
func (s PostService) Create(post *domain.Post) {
	s.PostSaver.Save(post)
}

// Delete post in application
func (s PostService) Delete(post domain.Post) {
	s.PostDeleter.Delete(post)
}
