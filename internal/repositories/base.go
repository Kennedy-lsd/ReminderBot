package repositories

import (
	"github.com/Kennedy-lsd/TelegramBot/databse"
	"gorm.io/gorm"
)

var DB *gorm.DB = databse.InitDB()
