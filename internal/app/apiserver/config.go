package apiserver

import "github.com/superdevdav/bash-app/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` // Адрес, на котором запускается сервер
	logLevel string `toml:"Log_level"`
	Store    *store.Config
}

// Инициализация config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		logLevel: "debug",
		Store:    store.NewConfig(),
	}
}
