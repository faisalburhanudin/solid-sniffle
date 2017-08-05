package service

import (
	"testing"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type mockForumSaver struct{}

func (mock mockForumSaver) Save(forum *domain.Forum){}

type mockForumNameChecker struct {
	isUsedReturn bool
}

func (mock mockForumNameChecker) IsUsed(forumName string) bool {
	return mock.isUsedReturn
}

func TestForumService_Create(t *testing.T) {
	forumService := ForumService{
		forumNameChecker: mockForumNameChecker{isUsedReturn:false},
		forumSaver: mockForumSaver{},
	}
	forumService.Create(&domain.Forum{})
}


func TestForumService_CreateNameUsed(t *testing.T) {
	forumService := ForumService{
		forumNameChecker: mockForumNameChecker{isUsedReturn:true},
		forumSaver: mockForumSaver{},
	}
	err := forumService.Create(&domain.Forum{})
	if err != ErrorForumNameUsed {
		t.Errorf("got: %v, want: %v.", err, ErrorForumNameUsed)
	}
}