package contract

import (
	"context"
	viewmodel "strider-challenge/app/viewmodel/homepage"
	"strider-challenge/domain/entity"
)

type HomeService interface {
	GetAllPosts(ctx context.Context, homePageRequestParams *viewmodel.HomePageRequest) (result entity.HomePageResponse, err error)
}
