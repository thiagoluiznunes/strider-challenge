package routehomepage

import (
	"strider-challenge/app/router/routeutils"
	viewmodel "strider-challenge/app/viewmodel/homepage"
	"strider-challenge/domain/contract"
	"strider-challenge/infra/config"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	cfg     *config.Config
	service contract.HomeService
}

func NewController(cfg *config.Config, service contract.HomeService) *Controller {
	return &Controller{
		cfg:     cfg,
		service: service,
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

	return routeutils.ResponseAPIOK(ctx, "OK")
}
