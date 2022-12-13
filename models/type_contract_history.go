package models

type ContractHistory struct {
	GormModel
	ContractID    uint   `json:"contractId"`
	UpdateBlock   uint64 `json:"updateBlock"`
	UpdateTime    uint64 `json:"updateTime"`
	UpdateTX      string `json:"updateTX"`
	PreviousOwner string `json:"previousOwner"`
	NewOwner      string `json:"newOwner"`
}
