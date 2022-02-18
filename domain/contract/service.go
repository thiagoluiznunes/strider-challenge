package contract

import (
	"context"
	"strider-challenge/app/viewmodel"
	"strider-challenge/domain/entity"
)

type HomeService interface {
	AddPost(ctx context.Context, postRequest *viewmodel.PostRequest) (err error)
	GetAllPosts(ctx context.Context, homePageRequestParams *viewmodel.HomePageRequest) (result entity.HomePageResponse, err error)
}

type UserService interface {
	GetUserByID(ctx context.Context, userRequest *viewmodel.UserRequest) (err error)
}
