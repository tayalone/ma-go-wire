package comment

import (
	"errors"

	CommentDomain "github.com/tayalone/ma-go-wire/comment/domain"
	Port "github.com/tayalone/ma-go-wire/comment/port"
)

type repo struct{}

// Initialize Comment Repository
func Initialize() Port.Repo {
	return &repo{}
}

func (r *repo) GetByPK(id uint) (CommentDomain.Domain, error) {
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

	return CommentDomain.Domain{}, errors.New("id does not exist")
	// / ------------------------
}

func (r *repo) GetByPostID(postID uint) ([]CommentDomain.Domain, error) {
	arr := []CommentDomain.Domain{}
	// /--- Mock Logic -----------
	for _, v := range Data {
		if v.PostID == postID {
			arr = append(arr, v)
		}
	}
	// / -------------------------
	if len(arr) > 0 {
		return arr, nil
	}
	return arr, errors.New(`post does not comments`)
}
