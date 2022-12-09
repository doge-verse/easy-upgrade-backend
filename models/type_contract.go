package models

// Contract .
type Contract struct {
	GormModel
	UserID             uint              `gorm:"NOT NULL" json:"userID"`
	Address            string            `gorm:"NOT NULL" json:"address"`
	AdminAddr          string            `gorm:"NOT NULL" json:"adminAddr"`
	Name               string            `json:"name"`
	Network            uint              `gorm:"NOT NULL" json:"network"`
	LastUpdate         uint64            `json:"lastUpdate"`
	Email              string            `json:"email"`
	ContractHistoryArr []ContractHistory `json:"contractHistoryArr,omitempty"`
}
