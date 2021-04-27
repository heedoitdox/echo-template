// Copyright Korea Space Data Co., Ltd. 2020 All Rights Reserved.
package config

import (
	"os"
	"sync"
)

type Configuration struct {
	DB_USERNAME  string `env:"DB_USERNAME"`
	DB_PASSWORD  string `env:"DB_PASSWORD"`
	DB_PORT      string `env:"DB_PORT"`
	DB_HOST      string `env:"DB_HOST"`
	DB_NAME      string `env:"DB_NAME"`
	TOKEN_SECRET string `env:"TOKEN_SECRET"`
}

var once sync.Once

func GetConfig() *Configuration {
	configuration := &Configuration{}
	once.Do(func() {
		configuration.DB_USERNAME = os.Getenv("DB_USERNAME")
		configuration.DB_PASSWORD = os.Getenv("DB_PASSWORD")
		configuration.DB_HOST = os.Getenv("DB_HOST")
		configuration.DB_PORT = os.Getenv("DB_PORT")
		configuration.DB_NAME = os.Getenv("DB_NAME")
		configuration.TOKEN_SECRET = os.Getenv("TOKEN_SECRET")
	})
	return configuration
}
