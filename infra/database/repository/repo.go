package repository

import (
	"database/sql"

	"github.com/thiagoluiznunes/strider-challenge/domain/contract"
	"github.com/thiagoluiznunes/strider-challenge/infra/database/repository/post"
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
