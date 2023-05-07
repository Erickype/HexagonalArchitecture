package repository

import (
	"errors"
	"fmt"
	"github.com/Erickype/HexagonalArchitecture/internal/core/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

func (m *MessengerPostgresRepository) SaveMessage(message domain.Message) error {
	request := m.db.Create(&message)
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

// NewMessengerPostgresRepository function that creates a MessengerPostgresRepository,
// creating a connection with postgres
func NewMessengerPostgresRepository() *MessengerPostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "Erickype"
	dbname := "hexagonal_architecture"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Message{})

	return &MessengerPostgresRepository{
		db: db,
	}
}
