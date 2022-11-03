package database

import (
	"github.com/Mikatech/CTF-AI/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Database() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}

	return db, err
}
