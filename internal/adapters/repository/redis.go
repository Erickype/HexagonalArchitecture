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
