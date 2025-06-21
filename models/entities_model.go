package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Fingerprint string    `gorm:"not null;"`
	CreatedAt   time.Time
	DeletedAt   *time.Time `gorm:"column:deleted_at"` // perhatikan penamaan kolom
}

type ChatHistory struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ImagePath string     `gorm:"type:text"`
	Question  string     `gorm:"type:text"`
	Answer    string     `gorm:"type:text"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// Gorm table name settings
func (Account) TableName() string     { return "account" }
func (ChatHistory) TableName() string { return "chat_history" }
