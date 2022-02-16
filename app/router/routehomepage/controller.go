package routehomepage

import (
	"strider-challenge/domain/contract"
	"strider-challenge/infra/config"
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
