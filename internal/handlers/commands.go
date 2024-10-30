package handlers

import (
	"github.com/Kennedy-lsd/TelegramBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "set_task":
		services.SetTask(bot, update)
	case "delete_task":
		services.DeleteTask(bot, update)
	case "show_all_tasks":
		services.ShowAllTasks(bot, update)
	}
}
