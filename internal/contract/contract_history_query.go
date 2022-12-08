package contract

import (
	"gorm.io/gorm"
)

// CHQuery .
type CHQuery struct {
	ContractAddr string
}

func (c CHQuery) chWhere() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.ContractAddr != "" {
			db = db.Where("contract_addr = ?", c.ContractAddr)
		}
		return db
	}
}
