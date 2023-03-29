package models

import (
	"time"
)

type Book struct{
	ID uint ` gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null; type:varchar(255)" json:"name_book"`
	Author string `gorm:"not null; type:varchar(255)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}