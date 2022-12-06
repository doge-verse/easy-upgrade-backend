package user

import (
	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
)

var (
	Repo repoI
)

func Init() {
	Repo = sqlRepo{
		db: sql.Db,
	}
}
