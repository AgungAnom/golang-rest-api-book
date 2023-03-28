package controllers

import (
	"gorm.io/gorm"
)

type Controllers struct {
	projectDB *gorm.DB
}

func New(db *gorm.DB) *Controllers {
	return &Controllers{
		projectDB: db,
	}
}