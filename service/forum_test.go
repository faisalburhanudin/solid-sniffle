package service

import (
	"testing"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type mockForumSaver struct{
	idChange int
}

func (mock mockForumSaver) Save(forum *domain.Forum){
	forum.Id = mock.idChange
}

type mockForumNameChecker struct {
	isUsedReturn bool
}

func (mock mockForumNameChecker) IsUsed(forumName string) bool {
	return mock.isUsedReturn
}

func TestForumService_Create(t *testing.T) {
	forumService := ForumService{
		forumNameChecker: mockForumNameChecker{isUsedReturn:false},
		forumSaver: mockForumSaver{idChange:1},
	}
	forum := domain.Forum{}
	forumService.Create(&forum)
	if forum.Id != 1 {
		t.Errorf("got: %v, want: %v.", forum.Id, 1)
	}
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