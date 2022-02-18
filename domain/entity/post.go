package entity

import (
	"strider-challenge/app/viewmodel"
	"strider-challenge/infra/exception"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        int64     `json:"id,omitempty" db:"id,omitempty"`
	UUID      *string   `json:"uuid,omitempty" db:"uuid,omitempty"`
	Type      string    `json:"type,omitempty" db:"type,omitempty"`
	Text      *string   `json:"text,omitempty" db:"text,omitempty"`
	UserID    *int64    `json:"user_id,omitempty" db:"user_id,omitempty"`
	PostID    *int64    `json:"post_id,omitempty" db:"post_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
}

func (p *Post) Validate() error {

	if p.UUID == nil || strings.TrimSpace(*p.UUID) == "" {
		return exception.NewValidationError("post_uuid", "empty/nil")
	}

	if strings.TrimSpace(p.Type) == "" {
		return exception.NewValidationError("type", "empty/nil")
	}

	if p.Text == nil || strings.TrimSpace(*p.Text) == "" {
		return exception.NewValidationError("text", "empty/nil")
	}

	if p.UserID == nil {
		return exception.NewValidationError("user_id", "empty/nil")
	}

	return nil
}

func BuilderPost(view *viewmodel.PostRequest) *Post {
	postUUID := uuid.New().String()
	return &Post{
		UUID:      &postUUID,
		Type:      view.Type,
		Text:      &view.Text,
		UserID:    &view.UserID,
		PostID:    &view.PostID,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}
