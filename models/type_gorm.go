package models

import (
	"time"
)

// GormModel base user
type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt int64      `json:"createdAt" sql:"index"`
	UpdatedAt int64      `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
