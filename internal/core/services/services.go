package services

import (
	"github.com/Erickype/HexagonalArchitecture/internal/core/domain"
	"github.com/Erickype/HexagonalArchitecture/internal/core/ports"
	"github.com/google/uuid"
)

type MessengerService struct {
	repository ports.MessengerRepository
}

func (m *MessengerService) SaveMessage(message domain.Message) error {
	message.Id = uuid.New().String()
	return m.repository.SaveMessage(message)
}

func (m *MessengerService) ReadMessage(id string) (*domain.Message, error) {
	return m.repository.ReadMessage(id)
}

func (m *MessengerService) ReadMessages() ([]*domain.Message, error) {
	return m.repository.ReadMessages()
}

// NewMessengerService creates and returns a new instance of MessengerService by passing a repository of type ports.MessengerRepository
func NewMessengerService(repository ports.MessengerRepository) *MessengerService {
	return &MessengerService{repository: repository}
}
