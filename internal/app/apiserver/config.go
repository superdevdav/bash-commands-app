package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"` // Адрес, на котором запускается сервер
	logLevel    string `toml:"Log_level"`
	DatabaseURL string `toml:"database_url"`
}

// Инициализация config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		logLevel: "debug",
	}
}
