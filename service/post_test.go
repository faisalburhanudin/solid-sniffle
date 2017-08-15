package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
	"testing"
)

type mockPostSaver struct{}

func (mock mockPostSaver) Save(post *domain.Post) {
	post.Id = 1
}

func TestPostService_Create(t *testing.T) {
	postService := PostService{PostSaver: mockPostSaver{}}
	post := domain.Post{}
	postService.Create(&post)
	if post.Id != 1 {
		t.Errorf("got: %v, want: %v.", post.Id, 1)
	}
}
