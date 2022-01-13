package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
