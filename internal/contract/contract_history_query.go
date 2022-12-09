package contract

import (
	"gorm.io/gorm"
)

// CHQuery .
type CHQuery struct {
	ContractID uint
}

func (c CHQuery) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.ContractID > 0 {
			db = db.Where("contract_id = ?", c.ContractID)
		}
		return db
	}
}
