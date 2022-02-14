package serverconfig

import (
	"strider-challenge/app/router/routeutils"

	"github.com/labstack/echo/v4"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(ctx echo.Context) error {
	return routeutils.ResponseAPIOK(ctx, "OK")
}
