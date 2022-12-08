package models

// Contract .
type Contract struct {
	GormModel
	UserID             uint   `gorm:"NOT NULL"`
	ContractAddr       string `gorm:"NOT NULL"`
	ContractName       string
	Network            uint `gorm:"NOT NULL"`
	LastUpdate         uint64
	Email              string
	ContractHistoryArr []ContractHistory
}
