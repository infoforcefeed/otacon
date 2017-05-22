package config

import "flag"

type Config struct {
	TrackerBind string `toml:"tracker_bind"`
	APIBind     string `toml:"api_bind"`
	DatabaseURL string `toml:"database_url"`
}

func New() (*Config, error) {
	c := Config{}

	flag.StringVar(&c.APIBind, "api-bind", "localhost:8080", "address to bind the api server to")
	flag.StringVar(&c.TrackerBind, "tracker-bind", "localhost:8081", "address to bind the tracker server to")
	flag.StringVar(&c.DatabaseURL, "database-url", "postgres://localhost:5432/otacon", "address to bind the api server to")

	flag.Parse()

	return &c, nil
}
