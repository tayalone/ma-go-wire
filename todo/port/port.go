package port

import TodoDomain "github.com/tayalone/ma-go-wire/todo/domain"

// Srv is Interface of Todo Service
type Srv interface {
	GetByID(id uint) (TodoDomain.Domain, error)
	GetByPostID(postID uint) ([]TodoDomain.Domain, error)
}

// Repo is Interface of Todo Repository
type Repo interface {
	GetByPK(id uint) (TodoDomain.Domain, error)
	GetByPostID(postID uint) ([]TodoDomain.Domain, error)
}
