package entity

import "time"

type Post struct {
	ID        int64     `json:"id,omitempty" db:"id,omitempty"`
	UUID      *string   `json:"uuid,omitempty" db:"uuid,omitempty"`
	Type      string    `json:"type,omitempty" db:"type,omitempty"`
	Text      *string   `json:"text,omitempty" db:"text,omitempty"`
	UserID    int64     `json:"user_id,omitempty" db:"user_id,omitempty"`
	PostID    *int64    `json:"post_id,omitempty" db:"post_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
}
