package databse

import (
	"fmt"
	"log"

	"github.com/Kennedy-lsd/TelegramBot/config"
	"github.com/Kennedy-lsd/TelegramBot/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbConfig := config.ConfigInit()

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		dbConfig.POSTGRES_HOST, dbConfig.POSTGRES_USER, dbConfig.POSTGRES_PASSWORD, dbConfig.POSTGRES_PORT, dbConfig.POSTGRES_DATABASE_NAME)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&data.Reminder{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return DB
}
