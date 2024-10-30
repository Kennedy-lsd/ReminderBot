package handlers

import (
	"fmt"

	"github.com/Kennedy-lsd/TelegramBot/internal/services"
	"github.com/Kennedy-lsd/TelegramBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, taskId := utils.GetKeyValue(update.CallbackQuery.Data)

	switch cmd {
	case "delete_task":
		if taskId == "" {
			sendErrorMessage(bot, update.CallbackQuery.Message.Chat.ID, "Invalid task ID.")
			return
		}
		services.DeleteTaskCallback(bot, update, taskId)
	default:
		sendErrorMessage(bot, update.CallbackQuery.Message.Chat.ID, "Unknown command.")
	}
}

func sendErrorMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(fmt.Sprintf("Failed to send error message: %v", err))
	}
}
