package main

import (
	"log"
	"os"
	"telegram-bot/config"
	"telegram-bot/internal/delivery/telegram"
	"telegram-bot/internal/repository"
	"telegram-bot/internal/usecase"
	"telegram-bot/pkg/telegram"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := telegram.NewClient(cfg.TelegramToken)
	repo := repository.NewMessageRepository()
	useCase := usecase.NewMessageUseCase(repo)
	handler := telegram.NewHandler(client, useCase)

	if err := handler.Start(); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}