package service

import (
	"github.com/faisalburhanudin/solid-sniffle/domain"
)

type PostSaver interface {
	Save(post *domain.Post)
}

type PostService struct {
	postSaver PostSaver
}

func (service PostService) Create(post *domain.Post) error {
	service.postSaver.Save(post)
	return nil
}
