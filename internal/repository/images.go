package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"astrolog/internal/model"
)

func (p Postgres) CreateImage(ctx context.Context, image model.Image) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.DB.ExecContext(queryCtx, "INSERT INTO images VALUES($1, $2, $3, $4, $5, $6)", image.ID, image.Copyright, image.Title, image.Explanation, image.URN, image.Date)
	if err != nil {
		return fmt.Errorf("exec context failed: %w", err)
	}

	return nil
}

func (p Postgres) GetImages(ctx context.Context) ([]model.Image, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := p.DB.QueryContext(queryCtx, "SELECT * FROM images")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNoImages
		}
		return nil, fmt.Errorf("query row context failed: %w", err)
	}

	var images []model.Image
	for rows.Next() {
		var image model.Image
		err := rows.Scan(&image.ID, &image.Copyright, &image.Title, &image.Explanation, &image.URN, &image.Date)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		images = append(images, image)
	}

	return images, nil
}

func (p Postgres) GetImageByDay(ctx context.Context, day time.Time) (model.Image, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := p.DB.QueryRowContext(queryCtx, "SELECT * FROM images WHERE date = $1", day)
	if row.Err() != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return model.Image{}, model.ErrNoImages
		}
		return model.Image{}, fmt.Errorf("query row context failed: %w", row.Err())
	}

	var image model.Image
	err := row.Scan(&image.ID, &image.Copyright, &image.Title, &image.Explanation, &image.URN, &image.Date)
	if err != nil {
		return model.Image{}, fmt.Errorf("scan failed: %w", err)
	}

	return image, nil
}
