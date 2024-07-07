package repository

import "telegram-bot/internal/domain"

type MessageRepository interface {
	Save(message *domain.Message) error
	Get(id int64) (*domain.Message, error)
}

type messageRepository struct {
	// Add any necessary fields for database connection
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) Save(message *domain.Message) error {
	// Implement save logic
	return nil
}

func (r *messageRepository) Get(id int64) (*domain.Message, error) {
	// Implement get logic
	return nil, nil
}
