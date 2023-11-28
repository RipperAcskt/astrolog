package model

import (
	"fmt"
	"time"
)

var (
	ErrNoImages = fmt.Errorf("error no images")
)

type Image struct {
	ID          string
	Copyright   string
	Explanation string
	Title       string
	URN         string
	Date        time.Time
}

type ImageApiResponse struct {
	Copyright   string `json:"copyright"`
	Explanation string `json:"explanation"`
	Title       string `json:"title"`
	URI         string `json:"hdurl"`
	Date        string `json:"date"`
}
