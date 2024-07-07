package telegram

import (
	"log"
	"telegram-bot/internal/domain"
	"telegram-bot/internal/usecase"
	"telegram-bot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	client  *telegram.Client
	useCase usecase.MessageUseCase
}

func NewHandler(client *telegram.Client, useCase usecase.MessageUseCase) *Handler {
	return &Handler{
		client:  client,
		useCase: useCase,
	}
}

func (h *Handler) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.client.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := &domain.Message{
			ID:   update.Message.Chat.ID,
			Text: update.Message.Text,
		}

		response, err := h.useCase.ProcessMessage(msg)
		if err != nil {
			log.Printf("Error processing message: %v", err)
			continue
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		_, err = h.client.Bot.Send(reply)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}

	return nil
}
