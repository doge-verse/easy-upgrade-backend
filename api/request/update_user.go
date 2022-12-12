package request

type UpdateEmail struct {
	UserID uint   `json:"-"`
	Email  string `json:"email"`
}
