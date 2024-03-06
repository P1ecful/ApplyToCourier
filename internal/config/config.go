package config

// configuration for Postgres
type PostgresConnection struct {
	Host     string
	Port     int
	Database string
	Password string
	Username string
}
