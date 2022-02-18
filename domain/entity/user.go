package entity

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	UUID      string    `json:"uuid,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserResponse struct {
	UserName          string    `json:"user_name"`
	NumberOfFollowers int64     `json:"number_of_followers"`
	NumberOfFollowing int64     `json:"number_of_following"`
	NumberOfPosts     int64     `json:"number_of_posts"`
	CreatedAt         time.Time `json:"created_at"`
}
