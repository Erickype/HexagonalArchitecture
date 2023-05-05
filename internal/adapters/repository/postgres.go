package repository

import (
	"errors"
	"fmt"
	"github.com/Erickype/HexagonalArchitecture/internal/core/domain"
	"github.com/jinzhu/gorm"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

func (m *MessengerPostgresRepository) SaveMessage(message domain.Message) error {
	request := m.db.Create(message)
	if request.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("message not saved: %t", request.Error))
	}
	return nil
}

func (m *MessengerPostgresRepository) ReadMessage(id string) (*domain.Message, error) {
	message := &domain.Message{}
	request := *m.db.First(&message, "id=?", id)
	if request.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("no message found: %t", request.Error))
	}
	return message, nil
}

func (m *MessengerPostgresRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	request := *m.db.Find(&messages)
	if request.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("no messages found: %t", request.Error))
	}
	return messages, nil
}
