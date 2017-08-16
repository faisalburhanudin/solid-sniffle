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

type PostsGetter interface {
	GetPosts() []domain.Post
}

// PostService using for inject dependencies via interface
type PostService struct {
	PostSaver   PostSaver
	PostDeleter PostDeleter
	PostsGetter PostsGetter
}

type PostDb interface {
	PostSaver
	PostDeleter
	PostsGetter
}

func NewPostService(post PostDb) *PostService {
	return &PostService{post, post, post}
}

// Create new post
func (s PostService) Create(post *domain.Post) {
	s.PostSaver.Save(post)
}

// Delete post in application
func (s PostService) Delete(post domain.Post) {
	s.PostDeleter.Delete(post)
}

// Gets list post
func (s PostService) GetPosts() []domain.Post {
	return s.PostsGetter.GetPosts()
}
