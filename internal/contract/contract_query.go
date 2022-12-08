package contract

import (
	"gorm.io/gorm"
)

// CQuery .
type CQuery struct {
	UserID  uint
	Address string
}

func (c CQuery) cWhere() func(db *gorm.DB) *gorm.DB {
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
