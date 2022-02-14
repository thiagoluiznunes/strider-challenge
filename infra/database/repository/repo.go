package repository

import (
	"database/sql"
	"strider-challenge/domain/contract"
)

type RepoManager struct {
	conn *sql.DB
}

func NewRepoManager(conn *sql.DB) contract.RepoManager {
	return &RepoManager{
		conn: conn,
	}
}
