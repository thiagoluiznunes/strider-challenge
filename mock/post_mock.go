package mock

import "github.com/google/uuid"

func NewUUID() *string {
	uuid := uuid.New().String()
	return &uuid
}

func NewString(s string) *string {
	return &s
}

func NewUserID(id int64) *int64 {
	return &id
}

const CreateDatabaseQuery = `CREATE DATABASE IF NOT EXISTS strider;`

const CreatePostsTableQuery = `
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
