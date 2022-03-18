package wire

import "github.com/google/uuid"

type FindByUser struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

type CreateList struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
