package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
)

var Repo repoI

func Init() {
	Repo = service{
		db: sql.Db,
	}
}
