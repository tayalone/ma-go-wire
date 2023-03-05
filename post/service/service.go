package service

import (
	PostDomain "github.com/tayalone/ma-go-wire/post/domain"
	Port "github.com/tayalone/ma-go-wire/post/port"
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

// GetByID return post by id or error
func (s *service) GetByID(id uint) (PostDomain.Domain, error) {
	return s.repo.GetByPK(id)
}
