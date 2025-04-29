package database

import (
	"fmt"
	"log"
	"luanti-skin-server/backend/models"
	"luanti-skin-server/backend/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	// Allow error logging only if debug enabled
	var dblogger = logger.Default.LogMode(logger.Silent)
	if utils.ConfigDebugDatabase {
		dblogger = logger.Default
	}

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("MT_SKIN_SERVER_DB_HOST"),
		os.Getenv("MT_SKIN_SERVER_DB_USER"),
		os.Getenv("MT_SKIN_SERVER_DB_PASSWORD"),
		os.Getenv("MT_SKIN_SERVER_DB_NAME"),
		os.Getenv("MT_SKIN_SERVER_DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dblogger,
	})

	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	err = DB.AutoMigrate(&models.Account{}, &models.Skin{})
	if err != nil {
		log.Fatalln("Failed to migrate database")
	}
}
