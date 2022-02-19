package routehomepage

import (
	"context"
	"net/http"

	"github.com/thiagoluiznunes/strider-challenge/app/router/routeutils"
	"github.com/thiagoluiznunes/strider-challenge/app/viewmodel"
	"github.com/thiagoluiznunes/strider-challenge/domain/contract"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	homeService contract.HomeService
}

func NewController(homeService contract.HomeService) *Controller {
	return &Controller{
		homeService: homeService,
	}
}

func (c *Controller) GetAllPosts(ctx echo.Context) error {

	homePageRequestParams := new(viewmodel.HomePageRequest)
	if err := homePageRequestParams.Binder(ctx); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	if err := homePageRequestParams.Validate(); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	result, err := c.homeService.GetAllPosts(context.Background(), homePageRequestParams)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, result)
}

func (c *Controller) AddPost(ctx echo.Context) error {

	postRequest := new(viewmodel.PostRequest)

	// Mocking UserID
	var userID int64 = 1
	postRequest.UserID = userID

	if err := ctx.Bind(postRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := postRequest.Validate(); err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	err := c.homeService.AddPost(context.Background(), postRequest)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, "OK")
}
