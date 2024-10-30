package data

import (
	"time"

	"github.com/google/uuid"
)

type Reminder struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	ChatId       int64     `gorm:"chat_id"`
	Message      string    `gorm:"type:text"`
	ReminderTime time.Time `gorm:"not null"`
}
