package config

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTg() *tgbotapi.BotAPI {
	tgConfig := ConfigInit()
	bot, err := tgbotapi.NewBotAPI(fmt.Sprintf("%v", tgConfig.TELEGRAM_APITOKEN))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}
