package routeuser

import (
	"strider-challenge/app/router/routeutils"
	"strider-challenge/app/viewmodel"
	"strider-challenge/domain/contract"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	userService contract.UserService
}

func NewController(userService contract.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (c *Controller) GetUser(ctx echo.Context) error {

	userRequestParams := new(viewmodel.HomePageRequest)
	if err := userRequestParams.Binder(ctx); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	if err := userRequestParams.Validate(); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	// result, err := c.userService.GetUserByID(context.Background(), userRequestParams)
	// if err != nil {
	// 	return routeutils.HandleAPIError(ctx, err)
	// }

	return routeutils.ResponseAPIOK(ctx, "K")
}
