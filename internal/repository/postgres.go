package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"astrolog/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	DB *sql.DB
}

func New(cfg config.Config) (Postgres, error) {
	db, err := sql.Open("pgx", cfg.GetDBUrl())
	if err != nil {
		return Postgres{}, fmt.Errorf("open failed: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return Postgres{}, fmt.Errorf("ping failed: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return Postgres{}, fmt.Errorf("with instance failed: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(cfg.MigratePath, "postgres", driver)
	if err != nil {
		return Postgres{}, fmt.Errorf("new with database instance failed: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return Postgres{}, fmt.Errorf("migration up failed: %w", err)
	}

	return Postgres{
		DB: db,
	}, nil
}

func (p Postgres) Close() error {
	err := p.DB.Close()
	return err
}
