package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"yujian-backend/pkg/log"
	"yujian-backend/pkg/model"
)

func InitDB(config model.DBConfig) {
	db := createConnect(config)
	userRepository = UserRepository{DB: db}
	postRepository = PostRepository{DB: db}
	bookRepository = BookRepository{DB: db}
}

func createConnect(config model.DBConfig) *gorm.DB {
	logger := log.GetLogger()
	db, err := gorm.Open(mysql.Open(config.CreateDsn()), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect database: %s", err)
		return nil
	} else {
		logger.Info("Connected to database")
		return db
	}
}
