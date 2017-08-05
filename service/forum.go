package service

import (
	"errors"
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type ForumNameChecker interface {
	IsUsed(forumName string) bool
}

type ForumSaver interface {
	Save(forum *domain.Forum)
}

type ForumService struct {
	forumNameChecker ForumNameChecker
	forumSaver       ForumSaver
}

var ErrorForumNameUsed = errors.New("forum name used")

// Create new form
func (service ForumService) Create(forum *domain.Forum) error {
	// Check forum name used
	isNameUsed := service.forumNameChecker.IsUsed(forum.Name)
	if isNameUsed {
		return ErrorForumNameUsed
	}

	// Save
	service.forumSaver.Save(forum)

	return nil
}
