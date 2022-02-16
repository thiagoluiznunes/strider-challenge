package routehomepage

import (
	"context"
	"strider-challenge/app/router/routeutils"
	viewmodel "strider-challenge/app/viewmodel/homepage"
	"strider-challenge/domain/contract"

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
