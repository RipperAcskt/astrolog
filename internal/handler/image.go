package handler

import (
	"errors"
	"fmt"
	"time"

	"astrolog/internal/model"

	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetImagesInfo(c *gin.Context) {
	imageInfo, err := h.s.GetImages(c.Request.Context())
	if err != nil {
		if errors.Is(err, model.ErrNoImages) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("get images info failed: %w", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, imageInfo)
}

func (h Handler) GetImageInfoByDay(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := time.Parse("2006-01-02", dayStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("date is not valid: %w", err).Error(),
		})
		return
	}

	imageInfo, err := h.s.GetImageByDay(c.Request.Context(), day)
	if err != nil {
		if errors.Is(err, model.ErrNoImages) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("get image info failed: %w", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, imageInfo)
}

func (h Handler) GetContent(context *gin.Context) {
	id := context.Param("id")
	context.File(fmt.Sprintf("store/%s", id))
}
