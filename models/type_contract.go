package models

// Contract .
type Contract struct {
	GormModel
	UserID             uint              `gorm:"NOT NULL" json:"userID"`
	ProxyAddress       string            `gorm:"NOT NULL;uniqueIndex:idx_uni;type:varchar(64)" json:"proxyAddress"`
	ProxyOwner         string            `gorm:"NOT NULL" json:"proxyOwner"`
	Name               string            `json:"name"`
	Network            uint              `gorm:"NOT NULL;uniqueIndex:idx_uni;type:varchar(64)" json:"network"` // chain id
	LastUpdate         uint64            `json:"lastUpdate"`
	Email              string            `json:"email"`
	ContractHistoryArr []ContractHistory `json:"contractHistoryArr,omitempty"`
}
