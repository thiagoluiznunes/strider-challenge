package entity

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	UUID      *string   `json:"uuid,omitempty"`
	UserName  *string   `json:"user_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
