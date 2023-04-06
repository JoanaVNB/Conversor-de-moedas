package repository

import (
	"exchange/domain"

	"gorm.io/gorm"
)

type Driver = *gorm.DB

type Repository struct{
	DB Driver
	err error
}

func NewRepository (DB Driver, err error) *Repository{
	return &Repository{DB: DB, err: err}
}

func Delete (DB Driver) {
	var lastExchange domain.Exchange
	DB.Last(&lastExchange)
	DB.Delete(&lastExchange)
}