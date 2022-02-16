package post

import "github.com/google/uuid"

func NewUUID() *string {
	uuid := uuid.New().String()
	return &uuid
}

func NewString(s string) *string {
	return &s
}

const CreateDatabaseQuery = `CREATE DATABASE IF NOT EXISTS strider;`

const CreatePostTableQuery = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER AUTO_INCREMENT NOT NULL,
		uuid VARCHAR(36) NOT NULL,
		type ENUM('original', 'repost', 'quote') NOT NULL,
		text VARCHAR(777) NOT NULL,
		user_id INTEGER NOT NULL,
		post_id INTEGER,
		updated_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (post_id) REFERENCES posts(id)
	) ENGINE = INNODB;
`
