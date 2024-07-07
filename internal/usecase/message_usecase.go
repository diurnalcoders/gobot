package usecase

import "telegram-bot/internal/domain"

type MessageUseCase interface {
	ProcessMessage(message *domain.Message) (string, error)
}

type messageUseCase struct {
	repo MessageRepository
}

func NewMessageUseCase(repo MessageRepository) MessageUseCase {
	return &messageUseCase{repo: repo}
}

func (uc *messageUseCase) ProcessMessage(message *domain.Message) (string, error) {
	// Implement your message processing logic here
	return "Processed: " + message.Text, nil
}
