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

// GetByID return comment or error
func (s *service) GetByID(id uint) (TodoDomain.Domain, error) {
	return s.repo.GetByPK(id)
}

// GetByPostID return comments or error
func (s *service) GetByPostID(postID uint) ([]TodoDomain.Domain, error) {
	return s.repo.GetByPostID(postID)
}
