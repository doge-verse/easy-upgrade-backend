package pkg

// Contract .
type Contract struct {
	GormModel
	UserID                  uint                   `json:"userID" gorm:"size:11"`
	ContractAddr            string                 `json:"contractAddr,omitempty"`
	ContractName            string                 `json:"contractName"`
	Network                 string                 `json:"network"`
	LastUpdate              int64                  `json:"lastUpdate"`
	ContractUpdateRecordArr []ContractUpdateRecord `json:"contractUpdateRecordArr,omitempty"`
}

type ContractUpdateRecord struct {
	GormModel
	ContractID  uint
	UpdateBlock uint
	UpdateTime  int64
}
