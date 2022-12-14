package subscriber

import (
	"log"

	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
)

func Init() {
	s := &Subscriber{
		Db: sql.Db,
	}

	contracts, err := s.SelectAllContract()
	if err != nil {
		log.Fatalln(err)
	}

	s.SubscribeAllContract(contracts)
}
