package port

import PostDomain "github.com/tayalone/ma-go-wire/post/domain"

// Srv is interface of Post Service
type Srv interface {
	GetByID(id uint) (PostDomain.Domain, error)
}

// Repo is interface of Post Repository
type Repo interface {
	GetByID(id uint) (PostDomain.Domain, error)
}
