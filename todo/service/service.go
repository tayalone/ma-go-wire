package service

import (
	TodoDomain "github.com/tayalone/ma-go-wire/todo/domain"
	Port "github.com/tayalone/ma-go-wire/todo/port"
)

type service struct {
	repo Port.Repo
}

// Initialize New Post Service
func Initialize(repo Port.Repo) Port.Srv {
	return &service{
		repo: repo,
	}
}

//
func (s *service) GetByID(id uint) (TodoDomain.Domain, error) {
	return s.repo.GetByPK(id)
}

func (s *service) GetByPostID(postID uint) ([]TodoDomain.Domain, error) {
	return s.repo.GetByPostID(postID)
}
