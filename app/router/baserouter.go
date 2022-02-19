package router

import (
	"github.com/thiagoluiznunes/strider-challenge/app/server/serverconfig"

	"github.com/labstack/echo/v4"
)

const (
	rootGroup = ""
	health    = "/health"
)

type Route struct {
}

func NewBaseRouter() *Route {
	return &Route{}
}

func (r *Route) Register(e *echo.Echo) {

	group := e.Group(rootGroup)
	group.GET(health, serverconfig.HealthCheck)
}
