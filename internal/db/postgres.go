package db

import (
	"applytocourier/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Repository struct {
	config *config.PostgresConnection
	log    *log.Logger
}

func NewRepository(cfg *config.PostgresConnection, log *log.Logger) *Repository {
	return &Repository{
		config: cfg,
		log:    log,
	}
}

func (r *Repository) Connect() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		r.config.Host, r.config.Port, r.config.Username, r.config.Password, r.config.Database))

	if err != nil {
		r.log.Fatal(err)
	}

	return db
}
