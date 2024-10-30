package repositories

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Kennedy-lsd/TelegramBot/data"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func SetTask(update tgbotapi.Update) error {
	pattern := `(.+) at (.+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(update.Message.Text)

	if len(matches) != 3 {
		return fmt.Errorf("invalid format. Use: 'Reminder: [Message] at [Time]'")
	}

	message := matches[1]
	reminderTimeStr := matches[2]

	reminderTime, err := time.Parse("02 Jan 15:04", reminderTimeStr)
	if err != nil {
		return fmt.Errorf("could not parse time: %v", err)
	}

	reminderTime = reminderTime.UTC().Truncate(time.Minute)

	task := data.Reminder{
		ID:           uuid.New(),
		ChatId:       update.Message.Chat.ID,
		Message:      message,
		ReminderTime: reminderTime,
	}

	if result := DB.Create(&task); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTask(taskId string) error {
	if result := DB.Where("id = ?", taskId).Delete(&data.Reminder{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllReminders(chatId int64) ([]data.Reminder, error) {
	var tasks []data.Reminder
	if result := DB.Where("chat_id = ?", chatId).Find(&tasks); result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}
