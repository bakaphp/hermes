package models

import "time"

// UsersFollows Model
type Messages struct {
	ID             uint
	AppsID         int `gorm:"size:11"`
	CompaniesID    int `gorm:"size:11"`
	UsersID        int `gorm:"size:11"`
	MessageTypesID int `gorm:"size:11"`
	Message        string
	ReactionsCount int `gorm:"size:11"`
	CommentsCount  int `gorm:"size:11"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	IsDeleted      int `gorm:"size:1"`
}
