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

	switch homePageRequestParams.Switch {
	case "all":
		result.Posts, err = s.svc.Repo.Post().GetAllPosts(ctx)
		if err != nil {
			return result, err
		}
	case "following":
		usersFollowingIDs := []interface{}{1, 2, 3}
		result.Posts, err = s.svc.Repo.Post().GetAllPostsByFollowing(ctx, usersFollowingIDs)
		if err != nil {
			return result, err
		}

	}

	return result, nil
}
