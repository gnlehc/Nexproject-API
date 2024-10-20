package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageID   uuid.UUID `json:"MessageID" gorm:"type:uuid;primaryKey;not null"`
	SenderID    uuid.UUID `json:"SenderID"`
	RecepientID uuid.UUID `json:"RecepientID"`
	FullName    string    `json:"FullName"`
	Message     string    `json:"Message"`
	CreatedAt   time.Time `json:"CreatedAt"`
}
