package viewmodel

import (
	"strings"

	"strider-challenge/infra/exception"

	"github.com/labstack/echo/v4"
)

type HomePageRequest struct {
	Switch string `query:"switch,omitempty"`
}

type PostRequest struct {
	UserID int64  `json:"user_id,omitempty"`
	Type   string `json:"type,omitempty"`
	Text   string `json:"text,omitempty"`
	PostID int64  `json:"post_id,omitempty"`
}

func (m *HomePageRequest) Binder(ctx echo.Context) error {
	return echo.QueryParamsBinder(ctx).
		String("switch", &m.Switch).
		BindError()
}

func (m *HomePageRequest) Validate() error {

	if strings.TrimSpace(m.Switch) == "" {
		return exception.NewValidationError("switch", "empty/nil")
	}

	return nil
}

func (m *PostRequest) Validate() error {

	if m.UserID == 0 {
		return exception.NewValidationError("user-uuid", "empty/nil")
	}

	if strings.TrimSpace(m.Type) == "" {
		return exception.NewValidationError("type", "empty/nil")
	}

	if strings.TrimSpace(m.Text) == "" {
		return exception.NewValidationError("switch", "empty/nil")
	}

	return nil
}
