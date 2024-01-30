package config

type PostgresConnection struct {
	Host     string
	Port     int
	Database string
	Password string
	Username string
}
