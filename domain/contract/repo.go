package contract

import (
	"context"

	"github.com/thiagoluiznunes/strider-challenge/domain/entity"
)

type RepoManager interface {
	Post() PostRepo
}

type PostRepo interface {
	Add(ctx context.Context, post entity.Post) (postID int64, err error)
	AddRepostOrQuote(ctx context.Context, post entity.Post) (postID int64, err error)
	GetAllPosts(ctx context.Context) (posts []entity.Post, err error)
	GetAllPostsByFollowing(ctx context.Context, usersFollowingIDs []interface{}) (posts []entity.Post, err error)
}
