package models

import "time"

// GroupMessages Model
type GroupMessages struct {
	MessageID int `gorm:"size:11"`
	GroupID   int `gorm:"size:11"`
	CreatedAt *time.Time
	IsDeleted int `gorm:"size:1"`
}
