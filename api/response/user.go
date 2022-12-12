package response

import "github.com/doge-verse/easy-upgrade-backend/models"

type UserInfo struct {
	models.User
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
