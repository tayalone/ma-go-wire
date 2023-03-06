package post

import (
	"errors"

	PostDomain "github.com/tayalone/ma-go-wire/post/domain"
	Port "github.com/tayalone/ma-go-wire/post/port"
)

type repo struct{}

// Initialize Comment Repository
func Initialize() Port.Repo {
	return &repo{}
}

func (r *repo) GetByPK(id uint) (PostDomain.Domain, error) {
	// /--- Mock Logic -----------
	index := -1
	for i, v := range Data {
		if v.ID == id {
			index = i
			break
		}
	}
	if index > -1 {
		return Data[index], nil
	}

	return PostDomain.Domain{}, errors.New("id does not exist")
	// / ------------------------
}
