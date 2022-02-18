package viewmodel

import (
	"strider-challenge/infra/exception"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	UserID string `query:"user_id,omitempty"`
}

func (m *UserRequest) Binder(ctx echo.Context) error {
	return echo.QueryParamsBinder(ctx).
		String("switch", &m.UserID).
		BindError()
}

func (m *UserRequest) Validate() error {

	if strings.TrimSpace(m.UserID) == "" {
		return exception.NewValidationError("switch", "empty/nil")
	}

	return nil
}
