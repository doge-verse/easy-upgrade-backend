package models

// Contract .
type Contract struct {
	GormModel
	UserID             uint
	ContractAddr       string
	ContractName       string
	Network            string
	LastUpdate         int64
	ContractHistoryArr []ContractHistory
}
