package entity

import "time"

type Follow struct {
	ID            int64     `json:"id,omitempty"`
	UserFollowed  int64     `json:"user_followed,omitempty"`
	UserFollowing int64     `json:"user_following,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
