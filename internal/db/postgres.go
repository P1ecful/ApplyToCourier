package db

import (
	"applytocourier/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.PostgresConnection, logger *log.Logger) *sql.DB {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database))

	if err != nil {
		logger.Fatal(err)
	}

	return db
}
