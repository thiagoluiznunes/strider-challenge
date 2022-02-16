package routehomepage

import "github.com/labstack/echo/v4"

const (
	rootGroup = "/"
)

type Route struct {
	name string
	ctrl *Controller
}

func NewRoute(name string, ctrl *Controller) *Route {
	return &Route{
		name: name,
		ctrl: ctrl,
	}
}

func (r *Route) Register(e *echo.Echo) {

	group := e.Group(rootGroup + r.name)
	group.GET(":switch", r.ctrl.GetAllPosts)

}
