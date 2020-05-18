package models

import "time"

// UserMessages Model
type UserMessages struct {
	MessageID int `gorm:"size:11"`
	UsersID   int `gorm:"size:11"`
	CreatedAt *time.Time
	IsDeleted int `gorm:"size:1"`
}
