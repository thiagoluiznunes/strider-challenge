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

func (r *PostRepository) GetAllPosts(ctx context.Context) (posts []entity.Post, err error) {

	const query = `
		SELECT
			uuid,
			type,
			text,
			user_id,
			post_id,
			updated_at,
			created_at
		FROM
			posts
	`

	result, err := r.conn.Query(query)
	if err != nil {
		return posts, err
	}

	for result.Next() {
		var post entity.Post
		err = result.Scan(
			&post.UUID,
			&post.Type,
			&post.Text,
			&post.UserID,
			&post.PostID,
			&post.UpdatedAt,
			&post.CreatedAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) GetAllPostsByFollowing(ctx context.Context, usersFollowingIDs []interface{}) (posts []entity.Post, err error) {

	var query = `
		SELECT
			uuid,
			type,
			text,
			user_id,
			post_id,
			updated_at,
			created_at
		FROM
			posts
		WHERE
			user_id IS NOT NULL
	`
	if len(usersFollowingIDs) > 0 {
		query += `AND user_id IN (`
		for index := range usersFollowingIDs {
			if index == len(usersFollowingIDs)-1 {
				query += `?`
				break
			}
			query += `?,`
		}
		query += `)`
	}

	result, err := r.conn.Query(query, usersFollowingIDs...)
	if err != nil {
		return posts, err
	}
	for result.Next() {
		var post entity.Post
		err = result.Scan(
			&post.UUID,
			&post.Type,
			&post.Text,
			&post.UserID,
			&post.PostID,
			&post.UpdatedAt,
			&post.CreatedAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
