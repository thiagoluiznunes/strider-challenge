package service

import (
	"github.com/thiagoluiznunes/strider-challenge/domain/contract"
	"github.com/thiagoluiznunes/strider-challenge/infra/config"
)

type Service struct {
	Config *config.Config
	Repo   contract.RepoManager
}

func NewService(cfg *config.Config, repo contract.RepoManager) *Service {
	return &Service{
		Config: cfg,
		Repo:   repo,
	}
}
