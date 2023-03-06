package port

import (
	PostDomain "github.com/tayalone/ma-go-wire/post/domain"
	TodoDomain "github.com/tayalone/ma-go-wire/todo/domain"
)

// Srv is interface of Post Service
type Srv interface {
	GetByID(id uint) (PostDomain.Domain, error)
	GetComments(id uint) ([]TodoDomain.Domain, error)
	GetComment(commentID uint) (TodoDomain.Domain, error)
}

// Repo is interface of Post Repository
type Repo interface {
	GetByPK(id uint) (PostDomain.Domain, error)
}
