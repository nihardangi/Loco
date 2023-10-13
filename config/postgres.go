package config

/*
PostgresConfig type struct
	contains address and whether postgres is enabled or not
*/
type PostgresConfig struct {
	URL      string
	User     string
	Password string
	Database string
	Port     int
}
