package models

import "time"

// GroupMessages Model
type GroupMessages struct {
	messageID int `gorm:"varchar(36)"`
	usersID   int `gorm:"size:11"`
	createdAt *time.Time
	isDeleted int `gorm:"size:1"`
}
