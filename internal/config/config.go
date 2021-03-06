package config

import (
	"net/url"
	"os"
)

// Config for the app
type Config struct {
	Host     string
	Port     string
	Addr     string
	Postgres url.URL
}

// New application config struct created
func New() Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := host + ":" + port

	sslmode := os.Getenv("DB_SSL_MODE")
	if sslmode == "" {
		sslmode = "require"
	}

	pg := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
		),
		Host: os.Getenv("DB_ADDRESS"),
		Path: os.Getenv("DB_NAME"),
	}

	v := url.Values{}

	v.Set("sslmode", sslmode)

	pg.RawQuery = v.Encode()

	return Config{host, port, addr, pg}

}
