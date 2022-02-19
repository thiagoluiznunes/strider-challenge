package mock

import (
	"context"
	"time"

	"github.com/thiagoluiznunes/strider-challenge/app/viewmodel"
	"github.com/thiagoluiznunes/strider-challenge/domain/contract"
	"github.com/thiagoluiznunes/strider-challenge/domain/entity"
)

type MockService struct {
}

func NewMockService() *MockService {
	return &MockService{}
}

type MockHomeService struct {
	svc MockService
}

func NewMockHomePageService(svc MockService) contract.HomeService {
	return &MockHomeService{
		svc: svc,
	}
}

func (s *MockHomeService) GetAllPosts(ctx context.Context, homePageRequestParams *viewmodel.HomePageRequest) (entity.HomePageResponse, error) {

	return entity.HomePageResponse{
		Posts: []entity.Post{
			{
				ID:        0,
				UUID:      NewUUID(),
				Type:      "original",
				Text:      NewUUID(),
				UserID:    NewUserID(1),
				PostID:    new(int64),
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			{
				ID:        0,
				UUID:      NewUUID(),
				Type:      "quote",
				Text:      NewUUID(),
				UserID:    NewUserID(2),
				PostID:    new(int64),
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
		},
	}, nil
}

func (s *MockHomeService) AddPost(ctx context.Context, postRequest *viewmodel.PostRequest) error {
	return nil
}
