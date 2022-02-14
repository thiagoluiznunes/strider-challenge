package service

import "strider-challenge/domain/contract"

type Service struct {
	Repo contract.RepoManager
}

func NewService(repo contract.RepoManager) *Service {
	return &Service{
		Repo: repo,
	}
}
