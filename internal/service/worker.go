package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"astrolog/config"
	"astrolog/internal/model"

	"github.com/google/uuid"
)

type workerService struct {
	repo workerRepo

	client http.Client
	timer  *time.Ticker

	cfg config.Config
}

type workerRepo interface {
	CreateImage(ctx context.Context, image model.Image) error
}

func newWorkerService(ctx context.Context, repo workerRepo, cfg config.Config) workerService {
	client := http.Client{
		Timeout: time.Second * time.Duration(cfg.ClientTimeout),
	}
	return workerService{
		repo: repo,

		client: client,
		timer:  time.NewTicker(24 * time.Hour),

		cfg: cfg,
	}
}

func (s workerService) startWorker(ctx context.Context) {
	for {
		select {
		case <-s.timer.C:
			image, err := s.getImageInfo()
			if err != nil {
				log.Printf("get image failed: %v", err)
			}

			err = s.repo.CreateImage(ctx, image)
			if err != nil {
				log.Printf("create image failed: %v", err)
			}
		}
	}

}

func (s workerService) getImageInfo() (model.Image, error) {
	uri := fmt.Sprintf("%s?api_key=%s", s.cfg.ApiURI, s.cfg.ApiKey)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return model.Image{}, fmt.Errorf("new request failed: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return model.Image{}, fmt.Errorf("client do failed: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Image{}, fmt.Errorf("read all failed: %w", err)
	}

	var imageResponse model.ImageApiResponse
	err = json.Unmarshal(body, &imageResponse)
	if err != nil {
		return model.Image{}, fmt.Errorf("unmarshall failed: %w", err)
	}

	id := uuid.NewString()
	err = s.getImage(id, imageResponse.URI)
	if err != nil {
		return model.Image{}, fmt.Errorf("get image failed: %w", err)
	}

	day, err := time.Parse("2006-01-02", imageResponse.Date)
	if err != nil {
		return model.Image{}, fmt.Errorf("parse failed: %w", err)
	}

	image := model.Image{
		ID:          id,
		Copyright:   imageResponse.Copyright,
		Explanation: imageResponse.Explanation,
		Title:       imageResponse.Title,
		URN:         fmt.Sprintf("%s/images/media/%s.png", s.cfg.ServerHost, id),
		Date:        day,
	}

	return image, nil
}

func (s workerService) getImage(id, uri string) error {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		log.Printf("new request failed: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("client do failed: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read all failed: %v", err)
	}

	fileName := fmt.Sprintf("store/%s.png", id)
	err = os.WriteFile(fileName, body, 0666)
	if err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}

	return nil
}
