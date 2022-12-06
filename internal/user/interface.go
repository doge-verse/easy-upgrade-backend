package user

import (
	"github.com/doge-verse/easy-upgrade-backend/pkg"
)

type repoI interface {
	GetUserByQuery(query Query) (*pkg.User, error)

	UserRegister(user *pkg.User) (*pkg.User, error)

	UpdateEmail(userID uint, email string) error
}
