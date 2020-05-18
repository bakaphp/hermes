package models

import "time"

// Users Model
type Users struct {
	id         int    `gorm:"primary_key"`
	title      string `gorm:"type:varchar(45)"`
	url        string `gorm:"type:varchar(45)"`
	languageID int
	createdAt  *time.Time
	updatedAt  *time.Time
	isDeleted  int `gorm:"size:1"`
}
