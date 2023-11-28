package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"astrolog/config"
	"astrolog/internal/handler"
	"astrolog/internal/repository"
	"astrolog/internal/server"
	"astrolog/internal/service"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config new failed: %w", err)
	}

	postgres, err := repository.New(cfg)
	if err != nil {
		return fmt.Errorf("postgres new failed: %w", err)
	}
	defer postgres.Close()

	svc := service.New(postgres, cfg)
	h, err := handler.New(svc, cfg)
	if err != nil {
		return fmt.Errorf("handler new failed: %w", err)
	}

	srv := server.Server{}

	go func() {
		if err := srv.Run(h.InitRouters(), cfg); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("server run failed: %v", err)
			return
		}
	}()

	if err := srv.ShutDown(); err != nil {
		return fmt.Errorf("server shut down failed: %w", err)
	}
	return nil
}
