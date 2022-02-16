package contract

import (
	"context"
	"strider-challenge/domain/entity"
)

type RepoManager interface {
	Post() PostRepo
}

type PostRepo interface {
	Add(ctx context.Context, post entity.Post) (postID int64, err error)
}
