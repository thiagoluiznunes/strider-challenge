package repository

import "strider-challenge/domain/contract"

type RepoManager struct {
	conn interface{} // example: mongoConn *mongo.Database
}

func NewRepoManager(conn interface{}) contract.RepoManager {
	return &RepoManager{
		conn: conn,
	}
}
