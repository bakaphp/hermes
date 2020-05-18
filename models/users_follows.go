package models

import "time"

// UsersFollows Model
type UsersFollows struct {
	ID              uint
	UsersID         int `gorm:"size:11"`
	EntityID        int `gorm:"size:11"`
	EntityNamespace string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	IsDeleted       int `gorm:"size:1"`
}
