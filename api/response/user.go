package response

type UserInfo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
