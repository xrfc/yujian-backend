package db

import (
	"log"
	"sync"
	"yujian-backend/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		// todo
		createConnect(model.DBConfig{})
	})
	return db
}

func createConnect(config model.DBConfig) {
	var err error
	db, err = gorm.Open(mysql.Open(config.CreateDsn()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
}
