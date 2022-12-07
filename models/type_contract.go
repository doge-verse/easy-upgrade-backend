package models

// Contract .
type Contract struct {
	GormModel
	UserID             uint              `json:"userID" gorm:"size:11"`
	ContractAddr       string            `json:"contractAddr"`
	ContractName       string            `json:"contractName"`
	Network            string            `json:"network"`
	LastUpdate         int64             `json:"lastUpdate"`
	ContractHistoryArr []ContractHistory `json:"contractHistoryArr,omitempty"`
}
