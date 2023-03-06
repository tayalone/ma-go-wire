package port

import CommentDomain "github.com/tayalone/ma-go-wire/comment/domain"

// Srv is Interface of Todo Service
type Srv interface {
	GetByID(id uint) (CommentDomain.Domain, error)
	GetByPostID(postID uint) ([]CommentDomain.Domain, error)
}

// Repo is Interface of Todo Repository
type Repo interface {
	GetByPK(id uint) (CommentDomain.Domain, error)
	GetByPostID(postID uint) ([]CommentDomain.Domain, error)
}
