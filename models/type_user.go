package models

// User .
type User struct {
	GormModel
	Name        string
	Address     string `gorm:"NOT NULL"`
	Email       string
	ContractArr []Contract
}
