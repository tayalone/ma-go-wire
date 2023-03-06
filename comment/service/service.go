package service

import (
	CommentDomain "github.com/tayalone/ma-go-wire/comment/domain"
	Port "github.com/tayalone/ma-go-wire/comment/port"
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
func (s *service) GetByID(id uint) (CommentDomain.Domain, error) {
	return s.repo.GetByPK(id)
}

// GetByPostID return comments or error
func (s *service) GetByPostID(postID uint) ([]CommentDomain.Domain, error) {
	return s.repo.GetByPostID(postID)
}
