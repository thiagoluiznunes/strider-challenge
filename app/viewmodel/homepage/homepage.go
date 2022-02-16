package viewmodel

import (
	"strider-challenge/infra/exception"
	"strings"

	"github.com/labstack/echo/v4"
)

type HomePageRequest struct {
	Switch string `query:"switch,omitempty"`
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
