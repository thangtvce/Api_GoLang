package models

import (
	"github.com/google/uuid"
)

type NotificationUser struct {
	ID             uint `gorm:"primaryKey"`
	NotificationID uuid.UUID
	UserID         uuid.UUID
}
