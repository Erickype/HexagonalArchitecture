package repository

import "github.com/jinzhu/gorm"

type MessengerPostgresRepository struct {
	db *gorm.DB
}
