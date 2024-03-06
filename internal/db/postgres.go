package db

import (
	"applytocourier/internal/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.PostgresConnection) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database))

	if err != nil {
		return nil, err
	}

	return db, err
}

// !FIXME
func MakeMigrations() {
	migrate, err := migrate.New(
		"file://db/ApplyToCourier",
		"postgres://postgres:postgres@localhost:5432/example?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Up(); err != nil {
		log.Fatal(err)
	}
}
