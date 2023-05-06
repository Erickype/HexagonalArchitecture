package repository

import (
	"encoding/json"
	"github.com/Erickype/HexagonalArchitecture/internal/core/domain"
	"github.com/go-redis/redis/v7"
)

type MessengerRedisRepository struct {
	client *redis.Client
}

func (m *MessengerRedisRepository) SaveMessage(message domain.Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	m.client.HSet("messages", message.Id, data)
	return nil
}

func (m *MessengerRedisRepository) ReadMessage(id string) (*domain.Message, error) {
	value, err := m.client.HGet("messages", id).Result()
	if err != nil {
		return nil, err
	}
	message := &domain.Message{}
	err = json.Unmarshal([]byte(value), message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (m *MessengerRedisRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	values, err := m.client.HGetAll("messages").Result()
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		message := &domain.Message{}
		err = json.Unmarshal([]byte(value), message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
