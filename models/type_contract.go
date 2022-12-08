package models

// Contract .
type Contract struct {
	GormModel
	UserID             uint   `gorm:"NOT NULL"`
	ContractAddr       string `gorm:"NOT NULL"`
	ContractName       string
	Network            string `gorm:"NOT NULL"`
	LastUpdate         int64
	ContractHistoryArr []ContractHistory
}
