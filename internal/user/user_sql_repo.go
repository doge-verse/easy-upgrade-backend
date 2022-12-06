package user

import (
	"github.com/doge-verse/easy-upgrade-backend/pkg"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

// GetUserByQuery .
func (repo sqlRepo) GetUserByQuery(query Query) (*pkg.User, error) {
	user := &pkg.User{}
	if err := repo.db.Model(user).Scopes(query.where()).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UserRegister .
func (repo sqlRepo) UserRegister(user *pkg.User) (*pkg.User, error) {
	if err := repo.db.Model(&pkg.User{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo sqlRepo) UpdateEmail(userID uint, email string) error {
	if err := repo.db.Model(&pkg.User{
		GormModel: pkg.GormModel{
			ID: userID,
		},
	}).Update("email", email).Error; err != nil {
		return err
	}
	return nil
}

func (repo sqlRepo) count(query Query) (int64, error) {
	var count int64
	if err := repo.db.Model(&pkg.User{}).Scopes(query.where()).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
