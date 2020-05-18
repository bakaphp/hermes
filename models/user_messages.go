package models

import "time"

// UserMessages Model
type UserMessages struct {
	messageID int `gorm:"varchar(36)"`
	usersID   int `gorm:"size:11"`
	createdAt *time.Time
	isDeleted int `gorm:"size:1"`
}
