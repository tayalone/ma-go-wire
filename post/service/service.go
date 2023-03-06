package service

import (
	PostDomain "github.com/tayalone/ma-go-wire/post/domain"
	Port "github.com/tayalone/ma-go-wire/post/port"
	TodoDomain "github.com/tayalone/ma-go-wire/todo/domain"
	TodoPort "github.com/tayalone/ma-go-wire/todo/port"
)

type service struct {
	repo    Port.Repo
	todoSrv TodoPort.Srv
}

// Initialize New Post Service
func Initialize(repo Port.Repo, todoSrv TodoPort.Srv) Port.Srv {
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
func (s *service) GetComments(id uint) ([]TodoDomain.Domain, error) {
	return s.todoSrv.GetByPostID(id)
}

// GetComment teturn comments or error
func (s *service) GetComment(commentID uint) (TodoDomain.Domain, error) {
	return s.todoSrv.GetByID(commentID)
}
