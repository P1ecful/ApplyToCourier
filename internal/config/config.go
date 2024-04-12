package config

// configuration for Postgres
type PostgresConnection struct {
	Host     string
	Port     string
	Database string
	Password string
	Username string
}
