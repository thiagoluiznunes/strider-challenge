package service

import (
	"context"
	viewmodel "strider-challenge/app/viewmodel/homepage"
	"strider-challenge/domain/contract"
	"strider-challenge/domain/entity"
)

type HomeService struct {
	svc Service
}

func NewHomePageService(svc Service) contract.HomeService {
	return &HomeService{
		svc: svc,
	}
}

func (s *HomeService) GetAllPosts(ctx context.Context, homePageRequestParams *viewmodel.HomePageRequest) (result entity.HomePageResponse, err error) {

	return result, nil
}
