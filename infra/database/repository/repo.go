package repository

import (
	"database/sql"
	"strider-challenge/domain/contract"
	"strider-challenge/infra/database/repository/post"
)

type RepoManager struct {
	conn *sql.DB
}

func NewRepoManager(conn *sql.DB) contract.RepoManager {
	return &RepoManager{
		conn: conn,
	}
}

func (r *RepoManager) Post() contract.PostRepo {
	return post.NewPostRepository(r.conn)
}
