package models

type ContractHistory struct {
	GormModel
	ContractID   uint
	UpdateBlock  uint
	UpdateTime   int64
	Network      string
	OperatorAddr string
}
