package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"astrolog/internal/mocks"
	"astrolog/internal/model"
	"astrolog/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetImages(t *testing.T) {
	type mockBehavior func(s *mocks.MockRepo)
	test := []struct {
		name         string
		mockBehavior mockBehavior
		err          error
	}{
		{
			name: "get images",
			mockBehavior: func(s *mocks.MockRepo) {
				s.EXPECT().GetImages(context.Background()).Return(
					[]model.Image{
						{
							ID:          "1",
							Copyright:   "1",
							Explanation: "1",
							Title:       "1",
							URN:         "1",
							Date:        time.Time{},
						},
					},
					nil,
				)
			},
			err: nil,
		},
		{
			name: "get images failed",
			mockBehavior: func(s *mocks.MockRepo) {
				s.EXPECT().GetImages(context.Background()).Return(
					nil,
					fmt.Errorf("1"),
				)
			},
			err: fmt.Errorf("get images failed: %w", fmt.Errorf("1")),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockRepo(ctrl)
			imageService := service.NewImageService(repo)

			tt.mockBehavior(repo)

			svc := service.Service{
				ImageService: imageService,
			}

			_, err := svc.GetImages(context.Background())
			assert.Equal(t, err, tt.err)
		})
	}
}

func TestGetImageByDay(t *testing.T) {
	type mockBehavior func(s *mocks.MockRepo)
	test := []struct {
		name         string
		mockBehavior mockBehavior
		err          error
	}{
		{
			name: "get image by day",
			mockBehavior: func(s *mocks.MockRepo) {
				s.EXPECT().GetImageByDay(context.Background(), gomock.Any()).Return(
					model.Image{
						ID:          "1",
						Copyright:   "1",
						Explanation: "1",
						Title:       "1",
						URN:         "1",
						Date:        time.Time{},
					},
					nil,
				)
			},
			err: nil,
		},
		{
			name: "get image by day failed",
			mockBehavior: func(s *mocks.MockRepo) {
				s.EXPECT().GetImageByDay(context.Background(), gomock.Any()).Return(
					model.Image{},
					fmt.Errorf("1"),
				)
			},
			err: fmt.Errorf("get image by day failed: %w", fmt.Errorf("1")),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockRepo(ctrl)
			imageService := service.NewImageService(repo)

			tt.mockBehavior(repo)

			svc := service.Service{
				ImageService: imageService,
			}

			_, err := svc.GetImageByDay(context.Background(), time.Time{})
			assert.Equal(t, err, tt.err)
		})
	}
}
