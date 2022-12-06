package pkg

import (
	"time"
)

// GormModel base user
type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt" sql:"index"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
