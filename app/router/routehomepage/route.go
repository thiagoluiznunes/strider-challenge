package routehomepage

import "github.com/labstack/echo/v4"

const (
	rootGroup = "/"
)

type Route struct {
	name string
	ctrl *Controller
}

func NewHomePageRoute(name string, ctrl *Controller) *Route {
	return &Route{
		name: name,
		ctrl: ctrl,
	}
}

func (r *Route) Register(e *echo.Echo) {

}
