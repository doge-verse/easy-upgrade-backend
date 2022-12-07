package user

import (
	"github.com/doge-verse/easy-upgrade-backend/models"
)

type repoI interface {
	GetUserByQuery(query Query) (*models.User, error)

	UserRegister(user *models.User) (*models.User, error)

	UpdateEmail(userID uint, email string) error
}
