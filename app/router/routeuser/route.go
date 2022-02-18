package routeuser

import "github.com/labstack/echo/v4"

const (
	rootGroup    = "/"
	followUser   = "/follow"
	unfollowUser = "/unfollow"
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

	group.GET(":id", r.ctrl.GetUser)
	group.POST(followUser, r.ctrl.FollowUser)
	group.POST(unfollowUser, r.ctrl.UnfollowUser)

}
