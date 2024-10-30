package services

import (
	"fmt"
	"time"

	"github.com/Kennedy-lsd/TelegramBot/data"
	"github.com/Kennedy-lsd/TelegramBot/internal/repositories"
	"github.com/Kennedy-lsd/TelegramBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, here you can create tasks for reminder.")
	msg.ReplyMarkup = utils.CmdKeyboard()
	sendMessage(bot, msg)
}

func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, write your task in format: Message at 28 Oct 23:11.")
	sendMessage(bot, msg)
}

func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Task successfully created"
	if err := repositories.SetTask(update); err != nil {
		text = "Couldn't set task"
	} else {
		taskTime := update.Message.Text
		confirmMsg := fmt.Sprintf("Reminder set for: %s", taskTime)
		sendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, confirmMsg))
	}
	sendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, text))
}

func DeleteTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	tasks, err := repositories.GetAllReminders(update.Message.Chat.ID)
	if err != nil {
		sendErrorMessage(bot, update.Message.Chat.ID, "Couldn't fetch tasks")
		return
	}

	if len(tasks) == 0 {
		sendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "No tasks available to delete."))
		return
	}

	keyboard := generateDeleteKeyboard(tasks)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, select the task you want to delete")
	msg.ReplyMarkup = keyboard
	sendMessage(bot, msg)
}

func DeleteTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskId string) {
	if err := repositories.DeleteTask(taskId); err != nil {
		sendErrorMessage(bot, update.CallbackQuery.Message.Chat.ID, "Couldn't delete the task")
		return
	}
	sendMessage(bot, tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Task successfully deleted"))
}

func sendReminder(bot *tgbotapi.BotAPI, task data.Reminder) {
	messageText := fmt.Sprintf("🔔 Reminder: %s", task.Message)
	msg := tgbotapi.NewMessage(task.ChatId, messageText)
	if _, err := bot.Send(msg); err != nil {
		fmt.Printf("Failed to send reminder: %v\n", err)
	}
}

func StartReminderChecker(bot *tgbotapi.BotAPI, db *gorm.DB) {
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			var dueTasks []data.Reminder
			startTime := time.Now().UTC().Add(-1 * time.Minute)
			endTime := time.Now().UTC()

			if err := db.Where("reminder_time BETWEEN ? AND ?", startTime, endTime).Find(&dueTasks).Error; err != nil {
				fmt.Println("Error fetching due tasks:", err)
				continue
			}

			for _, task := range dueTasks {
				sendReminder(bot, task)

				if err := db.Delete(&task).Error; err != nil {
					fmt.Printf("Error deleting completed task: %v\n", err)
				}
			}
		}
	}()
}

func ShowAllTasks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	tasks, err := repositories.GetAllReminders(update.Message.Chat.ID)
	if err != nil {
		sendErrorMessage(bot, update.Message.Chat.ID, "Couldn't get tasks")
		return
	}

	if len(tasks) == 0 {
		sendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "No tasks available."))
		return
	}

	text := "Tasks:\n"
	for _, task := range tasks {
		utcTime := task.ReminderTime.UTC()
		text += task.Message + " - " + utcTime.Format("02 Jan 15:04") + "\n"
	}

	sendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, text))
}

// Helper Functions

func sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
	}
}

func sendErrorMessage(bot *tgbotapi.BotAPI, chatID int64, errText string) {
	sendMessage(bot, tgbotapi.NewMessage(chatID, errText))
}

func generateDeleteKeyboard(tasks []data.Reminder) tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	for _, task := range tasks {
		btn := tgbotapi.NewInlineKeyboardButtonData(task.Message, "delete_task="+task.ID.String())
		buttons = append(buttons, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(buttons); i += 2 {
		if i+1 < len(buttons) {
			rows = append(rows, tgbotapi.NewInlineKeyboardRow(buttons[i], buttons[i+1]))
		} else {
			rows = append(rows, tgbotapi.NewInlineKeyboardRow(buttons[i]))
		}
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}
