package models

type ContractHistory struct {
	GormModel
	ContractID    uint
	UpdateBlock   uint
	Network       uint
	UpdateTime    uint64
	UpdateTX      string
	PreviousOwner string
	NewOwner      string
}
