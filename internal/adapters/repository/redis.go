package repository

import "github.com/go-redis/redis/v7"

type MessengerRedisRepository struct {
	client *redis.Client
}
