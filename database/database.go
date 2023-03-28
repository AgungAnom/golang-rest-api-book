package database

import (
	"fmt"
	"golang-rest-api-book/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres1"
	password = "postgres1"
	dbname = "db_book"
)

var (
	db *gorm.DB
	err error
)
func RunDB() *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host,port,user,password,dbname)
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Book{})
	fmt.Println("Connected to Database")
	return db
}