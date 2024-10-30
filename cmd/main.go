package main

import (
	"fmt"
	"log"

	"github.com/Kennedy-lsd/TelegramBot/config"
	"github.com/Kennedy-lsd/TelegramBot/databse"
	"github.com/Kennedy-lsd/TelegramBot/internal/handlers"
	"github.com/Kennedy-lsd/TelegramBot/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	apiConfig := config.ConfigInit()

	api := echo.New()
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	bot := config.InitTg()
	handlers.Update(bot)

	go services.StartReminderChecker(bot, databse.InitDB())

	log.Fatal(api.Start(fmt.Sprintf(":%s", apiConfig.PORT)))
}
