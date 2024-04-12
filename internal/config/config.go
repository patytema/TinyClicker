package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string `env:"TINY_CLICKER_TOKEN,required"`
}

func New() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
