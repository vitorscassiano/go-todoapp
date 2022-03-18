package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ListID      uuid.UUID `json:"id" gorm:"type:uuid;primary_key;not null;"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Tasks       []Task    `json:"tasks" gorm:"foreignKey:ListRef"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserRef     uuid.UUID
}

func (list *List) BeforeCreate(scope *gorm.DB) (err error) {
	list.ListID = uuid.New()

	return
}
