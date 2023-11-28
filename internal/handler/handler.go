package handler

import (
	"context"
	"time"

	"astrolog/config"
	"astrolog/internal/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s   imageService
	cfg config.Config
}

type imageService interface {
	GetImages(ctx context.Context) ([]model.Image, error)
	GetImageByDay(ctx context.Context, day time.Time) (model.Image, error)
}

func New(s imageService, cfg config.Config) (*Handler, error) {
	return &Handler{
		s:   s,
		cfg: cfg}, nil
}

func (h Handler) InitRouters() *gin.Engine {
	router := gin.New()

	images := router.Group("/images")
	images.GET("/", h.GetImagesInfo)
	images.GET("/:day", h.GetImageInfoByDay)
	images.GET("/media/:id", h.GetContent)

	return router
}
