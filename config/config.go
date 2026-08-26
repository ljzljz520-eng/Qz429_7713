package config

import "os"

type Config struct {
	Path string
	Port string
}

func Load() Config {
	p := os.Getenv("LAW_DB")
	if p == "" {
		p = "lawarchive.db"
	}
	port := os.Getenv("LAW_PORT")
	if port == "" {
		port = "8080"
	}
	return Config{p, port}
}
func (c Config) Address() string { return ":" + c.Port }
