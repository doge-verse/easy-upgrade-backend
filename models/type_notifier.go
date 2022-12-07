package models

type Notifier struct {
	GormModel
	ContractID   int64
	ContractAddr string
	Email        string
	UserID       int64
}
