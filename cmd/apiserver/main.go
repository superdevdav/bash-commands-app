package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/superdevdav/bash-app/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
