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

type Service struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Type struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type CallInfo struct {
    ServiceName string    `json:"service_name"`
    TypeName    string    `json:"type_name"`
    ChatID      string    `json:"chat_id"`
    Coin        string    `json:"coin"`
    CreatedAt   time.Time `json:"created_at"`
}