package service

import (
	"testing"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type mockPostSaver struct{}

func (mock mockPostSaver) Save(post *domain.Post) {
	post.Id = 1
}

func TestPostService_Create(t *testing.T) {
	postService := PostService{postSaver: mockPostSaver{}}
	post := domain.Post{}
	postService.Create(&post)
	if post.Id != 1 {
		t.Errorf("got: %v, want: %v.", post.Id, 1)
	}
}
