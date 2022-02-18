package viewmodel

import (
	"strider-challenge/infra/exception"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	UserID          int64 `json:"user_id,omitempty"`
	FollowingUserID int64 `json:"following_user_id,omitempty"`
}

func (m *UserRequest) Binder(ctx echo.Context) error {
	err := echo.QueryParamsBinder(ctx).
		Int64("user_id", &m.UserID).
		BindError()
	if err != nil {
		return err
	}
	return nil
}

func (m *UserRequest) Validate() error {

	if m.UserID == 0 {
		return exception.NewValidationError("user_id", "empty/nil")
	}

	if m.UserID == m.FollowingUserID {
		return exception.NewValidationError("user_id", "users cannont follow/unfollow themselves")
	}

	return nil
}
