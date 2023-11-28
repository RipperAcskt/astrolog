package service

import (
	"context"
	"fmt"
	"time"

	"astrolog/internal/model"
)

type ImageService struct {
	repo imageRepo
}

type imageRepo interface {
	GetImages(ctx context.Context) ([]model.Image, error)
	GetImageByDay(ctx context.Context, day time.Time) (model.Image, error)
}

func NewImageService(repo imageRepo) ImageService {
	return ImageService{
		repo: repo,
	}
}

func (s ImageService) GetImages(ctx context.Context) ([]model.Image, error) {
	images, err := s.repo.GetImages(ctx)
	if err != nil {
		return nil, fmt.Errorf("get images failed: %w", err)
	}

	return images, nil
}

func (s ImageService) GetImageByDay(ctx context.Context, day time.Time) (model.Image, error) {
	image, err := s.repo.GetImageByDay(ctx, day)
	if err != nil {
		return model.Image{}, fmt.Errorf("get image by day failed: %w", err)
	}

	return image, nil
}
