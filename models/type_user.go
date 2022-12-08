package models

// User .
type User struct {
	GormModel
	Name        string
	Address     string
	Email       string
	ContractArr []Contract
}
