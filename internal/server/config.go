package server

import (
	"time"

	"github.com/nicholasjackson/env"
)

type Config struct {
	Addr         string
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DBClientURI  string
	DBName       string
}

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func NewConfig() *Config {
	env.Parse()

	return &Config{
		Addr:         *bindAddress,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		DBClientURI:  "mongodb://mongo:27017",
		DBName:       "snpt",
	}
}
