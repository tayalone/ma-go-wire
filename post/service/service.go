package service

import (
	CommentDomain "github.com/tayalone/ma-go-wire/comment/domain"
	CommentPort "github.com/tayalone/ma-go-wire/comment/port"
	PostDomain "github.com/tayalone/ma-go-wire/post/domain"
	Port "github.com/tayalone/ma-go-wire/post/port"
)

type service struct {
	repo    Port.Repo
	todoSrv CommentPort.Srv
}

// Initialize New Post Service
func Initialize(repo Port.Repo, todoSrv CommentPort.Srv) Port.Srv {
	return &service{
		repo:    repo,
		todoSrv: todoSrv,
	}
}

// GetByID return post by id or error
func (s *service) GetByID(id uint) (PostDomain.Domain, error) {
	return s.repo.GetByPK(id)
}

// GetComments teturn []comments or error
func (s *service) GetComments(id uint) ([]CommentDomain.Domain, error) {
	return s.todoSrv.GetByPostID(id)
}

// GetComment teturn comments or error
func (s *service) GetComment(commentID uint) (CommentDomain.Domain, error) {
	return s.todoSrv.GetByID(commentID)
}
