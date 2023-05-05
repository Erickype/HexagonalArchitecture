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
