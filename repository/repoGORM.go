package repository

import (
	"exchange/domain"
	"gorm.io/gorm"
)

type Repository struct{
	DB *gorm.DB
}

func NewRepository (DB *gorm.DB) *Repository{
	return &Repository{DB: DB}
}

func Delete (DB *gorm.DB) {
	var lastExchange domain.Exchange
	DB.Last(&lastExchange)
	DB.Delete(&lastExchange)
}