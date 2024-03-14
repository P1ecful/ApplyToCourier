package config

import (
	"os"
)

// configuration for Postgres
type PostgresConnection struct {
	Host     string `yaml:"POSTGRES_HOST"`
	Port     string `yaml:"POSTGRES_PORT"`
	Database string `yaml:"POSTGRES_DB"`
	Password string `yaml:"POSTGRES_PASSWORD"`
	Username string `yaml:"POSTGRES_USER"`
}

// getting meta to connect Database
func LoadPostgresEnv() *PostgresConnection {
	return &PostgresConnection{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Username: os.Getenv("POSTGRES_USER"),
	}
}
