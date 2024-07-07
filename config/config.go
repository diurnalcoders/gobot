package config

import "os"

type Config struct {
	TelegramToken string
}

func Load() (*Config, error) {
	return &Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}, nil
}
