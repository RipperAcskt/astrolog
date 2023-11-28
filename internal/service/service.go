package service

import (
	"context"

	"astrolog/config"
)

type Service struct {
	ImageService
	workerService
}

//go:generate mockgen -destination=../mocks/mock_repo.go -package=mocks astrolog/internal/service Repo
type Repo interface {
	imageRepo
	workerRepo
}

func New(repo Repo, cfg config.Config) Service {
	svc := Service{
		ImageService:  NewImageService(repo),
		workerService: newWorkerService(context.Background(), repo, cfg),
	}
	go svc.startWorker(context.Background())
	return svc
}
