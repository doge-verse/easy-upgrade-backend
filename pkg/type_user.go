package pkg

// User .
type User struct {
	GormModel
	Name        string     `json:"name" gorm:"size:100"`
	Address     string     `json:"address" gorm:"size:32"`
	Email       string     `json:"email"`
	ContractArr []Contract `json:"contractArr,omitempty"`
}
