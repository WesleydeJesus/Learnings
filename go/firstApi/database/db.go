package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	err error
)

func ConnectDB(){
	stgConnect := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stgConnect))
	if err != nil {
		log.Panic("Error Connect")
	}
}