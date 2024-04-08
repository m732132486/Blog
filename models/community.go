package models

import (
	"time"

	"gorm.io/gorm"
)

type Communities struct {
	CategoryID   int64  `json:"category_id"db:"category_id" binding:"required"`
	CategoryName string `json:"category_name"db:"category_name"`
}

func (table *Communities) TableName() string {
	return "communities"
}

type CommunityCategory struct {
	gorm.Model
	ID           int64  `json:"id"db:"id"`
	CategoryID   int64  `json:"category_id"db:"category_id"`
	CategoryName string `json:"category_name"db:"category_name"`
	ParentID     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
