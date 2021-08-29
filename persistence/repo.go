package persistence

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return nil, err
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
