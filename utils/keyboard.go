package utils

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdKeyboard() tgbotapi.ReplyKeyboardMarkup {
	var cmdKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/set_task"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/delete_task"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/show_all_tasks"),
		),
	)
	return cmdKeyboard
}

func GetKeyValue(data string) (string, string) {
	parts := strings.SplitN(data, "=", 2)
	if len(parts) < 2 {
		return data, ""
	}
	return parts[0], parts[1]
}
