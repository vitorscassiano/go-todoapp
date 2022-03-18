package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    uuid.UUID `json:"id" gorm:"type:uuid;promary_key;not null;"`
	Name      string    `json:"name"`
	Lists     []List    `json:"lists" gorm:"foreignKey:UserRef"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(scope *gorm.DB) (err error) {
	user.UserID = uuid.New()

	return
}
