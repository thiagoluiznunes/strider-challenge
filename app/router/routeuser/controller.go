package routeuser

import (
	"context"
	"net/http"
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

	userRequestParams := new(viewmodel.UserRequest)
	if err := userRequestParams.Binder(ctx); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	if err := userRequestParams.Validate(); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	result, err := c.userService.GetUserByID(context.Background(), userRequestParams)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, result)
}

func (c *Controller) FollowUser(ctx echo.Context) error {

	userRequest := new(viewmodel.UserRequest)

	// Mocking UserID
	var userID int64 = 1
	userRequest.UserID = userID

	if err := ctx.Bind(userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := userRequest.Validate(); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	err := c.userService.FollowUser(context.Background(), userRequest)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, "OK")
}

func (c *Controller) UnfollowUser(ctx echo.Context) error {

	userRequest := new(viewmodel.UserRequest)

	var userID int64 = 1
	userRequest.UserID = userID

	if err := ctx.Bind(userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := userRequest.Validate(); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	err := c.userService.UnfollowUser(context.Background(), userRequest)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, "OK")
}
