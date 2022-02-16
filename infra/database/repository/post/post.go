package post

import (
	"context"
	"database/sql"
	"strider-challenge/domain/contract"
	"strider-challenge/domain/entity"
	"strider-challenge/infra/exception"
)

type PostRepository struct {
	conn *sql.DB
}

func NewPostRepository(conn *sql.DB) contract.PostRepo {
	return &PostRepository{
		conn: conn,
	}
}

func (r *PostRepository) Add(ctx context.Context, post entity.Post) (postID int64, err error) {

	const query = `
		INSERT INTO posts (
			uuid,
			type,
			text,
			user_id,
			post_id,
			updated_at,
			created_at
		)
		VALUE (
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)
	`

	result, err := r.conn.Exec(query,
		post.UUID,
		post.Type,
		post.Text,
		post.UserID,
		post.PostID,
		post.UpdatedAt,
		post.CreatedAt,
	)
	if err != nil {
		return postID, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return postID, err
	} else if count <= 0 {
		return postID, exception.NewConflictError("not rows affected")
	}

	postID, err = result.LastInsertId()
	if err != nil {
		return postID, err
	}

	return postID, nil
}
