package port

import (
	CommentDomain "github.com/tayalone/ma-go-wire/comment/domain"
	PostDomain "github.com/tayalone/ma-go-wire/post/domain"
)

// Srv is interface of Post Service
type Srv interface {
	GetByID(id uint) (PostDomain.Domain, error)
	GetComments(id uint) ([]CommentDomain.Domain, error)
	GetComment(commentID uint) (CommentDomain.Domain, error)
}

// Repo is interface of Post Repository
type Repo interface {
	GetByPK(id uint) (PostDomain.Domain, error)
}
