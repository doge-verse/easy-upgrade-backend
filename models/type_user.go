package models

// User .
type User struct {
	GormModel
	Name        string     `json:"name"`
	Address     string     `gorm:"NOT NULL;uniqueIndex:idx_uni;type:varchar(64)" json:"address"`
	Email       string     `json:"email"`
	ContractArr []Contract `json:"contractArr,omitempty"`
}
