package service

import (
	"context"
	"time"

	"strider-challenge/app/viewmodel"
	"strider-challenge/domain/contract"
	"strider-challenge/domain/entity"
	"strider-challenge/infra/exception"
)

var MockUser entity.User
var MockUserResponse entity.UserResponse
var MockFollowedUsersIDs []int64

type UserService struct {
	svc Service
}

func NewUserService(svc Service) contract.UserService {
	MockUser = entity.User{
		ID:        1,
		UUID:      "xxxx-xxxxx-xxxxx-xxxxx",
		UserName:  "thiagonunes",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	MockUserResponse = entity.UserResponse{
		UserName:          MockUser.UserName,
		NumberOfFollowers: 0,
		NumberOfFollowing: int64(len(MockFollowedUsersIDs)),
		NumberOfPosts:     0,
		CreatedAt:         MockUser.CreatedAt,
	}
	return &UserService{
		svc: svc,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, userReq *viewmodel.UserRequest) (entity.UserResponse, error) {
	return MockUserResponse, nil
}

func (s *UserService) FollowUser(ctx context.Context, userReq *viewmodel.UserRequest) error {

	for _, value := range MockFollowedUsersIDs {
		if value == userReq.FollowingUserID {
			return exception.NewConflictError("user already followed")
		}
	}
	MockFollowedUsersIDs = append(MockFollowedUsersIDs, userReq.FollowingUserID)
	MockUserResponse.NumberOfPosts = int64(len(MockFollowedUsersIDs))

	return nil
}

func (s *UserService) UnfollowUser(ctx context.Context, userReq *viewmodel.UserRequest) error {

	for index, value := range MockFollowedUsersIDs {
		if value == userReq.FollowingUserID {
			MockFollowedUsersIDs = append(MockFollowedUsersIDs[:index], MockFollowedUsersIDs[index+1:]...)
			MockUserResponse.NumberOfPosts = int64(len(MockFollowedUsersIDs))
			break
		}
	}

	return nil
}
