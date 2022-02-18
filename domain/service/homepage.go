package service

import (
	"context"
	"strider-challenge/app/viewmodel"
	"strider-challenge/domain/contract"
	"strider-challenge/domain/entity"
	"strider-challenge/infra/exception"
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
			return result, exception.NewApplicationError(err)
		}
	case "following":
		// Mocked users ids
		usersFollowingIDs := []interface{}{1, 2, 3}
		result.Posts, err = s.svc.Repo.Post().GetAllPostsByFollowing(ctx, usersFollowingIDs)
		if err != nil {
			return result, exception.NewApplicationError(err)
		}
	}

	return result, nil
}

func (s *HomeService) AddPost(ctx context.Context, postRequest *viewmodel.PostRequest) (err error) {

	post := entity.BuilderPost(postRequest)
	err = post.Validate()
	if err != nil {
		return exception.NewApplicationError(err)
	}

	switch postRequest.Type {
	case "original":
		_, err = s.svc.Repo.Post().Add(ctx, *post)
		if err != nil {
			return exception.NewApplicationError(err)
		}
	default:
		var mockRepostID int64 = 1
		post.PostID = &mockRepostID
		_, err = s.svc.Repo.Post().AddRepostOrQuote(ctx, *post)
		if err != nil {
			return exception.NewApplicationError(err)
		}
	}

	return nil
}
