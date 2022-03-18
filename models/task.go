package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskID      uuid.UUID `json:"id" gorm:"type:uuid;primary_key;not null;"`
	Icon        string    `json:"icon"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ListRef     uuid.UUID
}

func (task *Task) BeforeCreate(scope *gorm.DB) (err error) {
	task.TaskID = uuid.New()

	return
}
