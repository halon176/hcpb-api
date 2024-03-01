package db

import (
	"time"

	"github.com/google/uuid"
)

type Call struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ServiceID uint      `json:"service_id"`
	TypeID    uint      `json:"type_id"`
	ChatID    string    `json:"chat_id"`
	Coin      string    `json:"coin"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW"`
}
