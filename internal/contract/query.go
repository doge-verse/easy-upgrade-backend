package contract

import (
	"gorm.io/gorm"
)

// Query .
type Query struct {
	Limit   uint
	Offset  uint
	UserID  uint
	Address string
}

func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Address != "" {
			db = db.Where("address = ?", c.Address)
		}
		if c.UserID > 0 {
			db = db.Where("user_id = ?", c.UserID)
		}
		return db
	}
}
