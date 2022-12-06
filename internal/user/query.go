package user

import (
	"gorm.io/gorm"
)

// Query .
type Query struct {
	Limit   uint
	Offset  uint
	UserID  uint
	Email   string
	Address string
}

// Where
func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.UserID > 0 {
			db = db.Where("id = ?", c.UserID)
		}
		if len(c.Email) > 0 {
			db = db.Where("email = ?", c.Email)
		}
		if len(c.Address) > 0 {
			db = db.Where("address = ?", c.Address)
		}
		db = db.Preload("ContractArr")
		return db
	}
}
