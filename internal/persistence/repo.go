package persistence

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dbURL string) (*Repository, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	err = makeMigrations(db)
	if err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("making migration: %w", err)
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Close() {
	if err := r.db.Close(); err != nil {
		log.Fatalln(err)
	}
}

func makeMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	return m.Up()
}
