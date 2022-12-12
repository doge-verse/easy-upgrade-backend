package contract

import (
	"gorm.io/gorm"
)

// CQuery .
type CQuery struct {
	UserID       uint
	ProxyAddress string
}

func (c CQuery) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.ProxyAddress != "" {
			db = db.Where("proxy_address = ?", c.ProxyAddress)
		}
		if c.UserID > 0 {
			db = db.Where("user_id = ?", c.UserID)
		}
		return db
	}
}
